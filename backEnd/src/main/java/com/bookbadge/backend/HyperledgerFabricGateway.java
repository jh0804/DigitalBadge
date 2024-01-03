package com.bookbadge.backend;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonParser;
import io.grpc.Grpc;
import io.grpc.ManagedChannel;
import io.grpc.TlsChannelCredentials;
import org.hyperledger.fabric.client.CommitException;
import org.hyperledger.fabric.client.CommitStatusException;
import org.hyperledger.fabric.client.Contract;
import org.hyperledger.fabric.client.EndorseException;
import org.hyperledger.fabric.client.Gateway;
import org.hyperledger.fabric.client.GatewayException;
import org.hyperledger.fabric.client.SubmitException;
import org.hyperledger.fabric.client.identity.Identities;
import org.hyperledger.fabric.client.identity.Identity;
import org.hyperledger.fabric.client.identity.Signer;
import org.hyperledger.fabric.client.identity.Signers;
import org.hyperledger.fabric.client.identity.X509Identity;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.cert.CertificateException;
import java.time.Instant;
import java.util.concurrent.TimeUnit;

public class HyperledgerFabricGateway {

    private static final String CHANNEL_NAME = System.getenv().getOrDefault("CHANNEL_NAME", "mychannel");
    private static final String CHAINCODE_NAME = System.getenv().getOrDefault("CHAINCODE_NAME", "basic");

    public static ManagedChannel newGrpcConnection(String peerEndpoint, Path tlsCertPath,String overrideAuth) throws IOException {
        var credentials = TlsChannelCredentials.newBuilder()
                .trustManager(tlsCertPath.toFile())
                .build();
        return Grpc.newChannelBuilder(peerEndpoint, credentials)
                .overrideAuthority(overrideAuth)
                .build();
    }

    private static Identity newIdentity(Path certPath, String mspId) throws IOException, CertificateException {
        var certReader = Files.newBufferedReader(certPath);
        var certificate = Identities.readX509Certificate(certReader);
        return new X509Identity(mspId, certificate);
    }

    private static Signer newSigner(Path keyDirPath) throws IOException, InvalidKeyException {
        try (var keyFiles = Files.list(keyDirPath)) {
            var privateKeyPath = keyFiles.findFirst().orElseThrow();
            var keyReader = Files.newBufferedReader(privateKeyPath);
            var privateKey = Identities.readPrivateKey(keyReader);
            return Signers.newPrivateKeySigner(privateKey);
        }
    }

    public static Contract initializeContract(
            String mspId,
            Path tlsCertPath,
            Path cryptoPath,
            Path keyDirPath,
            Path certPath,
            String peerEndpoint,
            String overrideAuth
    ) {
        try {
            var channel = newGrpcConnection(peerEndpoint, tlsCertPath, overrideAuth);
            var identity = newIdentity(certPath, mspId);
            var signer = newSigner(keyDirPath);

            var builder = Gateway.newInstance().identity(identity).signer(signer).connection(channel)
                    .evaluateOptions(options -> options.withDeadlineAfter(5, TimeUnit.SECONDS))
                    .endorseOptions(options -> options.withDeadlineAfter(15, TimeUnit.SECONDS))
                    .submitOptions(options -> options.withDeadlineAfter(5, TimeUnit.SECONDS))
                    .commitStatusOptions(options -> options.withDeadlineAfter(1, TimeUnit.MINUTES));

            var gateway = builder.connect();
            var network = gateway.getNetwork(CHANNEL_NAME);
            return network.getContract(CHAINCODE_NAME); // = contract
        } catch (Exception e) {
            throw new RuntimeException("Failed to initialize contract", e);
        }
    }
}


