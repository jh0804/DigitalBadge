package com.bookbadge.backend;

import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import com.bookbadge.backend.badge.Badge;
import com.bookbadge.backend.badge.BadgeResponseDto;
import com.bookbadge.backend.bogo.Bogo;
import com.bookbadge.backend.bogo.BogoRequestDto;
import com.bookbadge.backend.bogo.BogoResponseDto;
import com.bookbadge.backend.member.Lib;
import com.bookbadge.backend.member.LibRepository;
import com.bookbadge.backend.member.MemDto;
import com.bookbadge.backend.member.Member;
import com.bookbadge.backend.member.Stu;
import com.bookbadge.backend.member.StuRepository;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;

import java.time.Instant;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

import javax.annotation.PostConstruct;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;

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

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.Arrays;



@Service
public class LibraryService {

    @Autowired
    private LibRepository libRepository;

    String MSP_ID = System.getenv().getOrDefault("MSP_ID", "Org2MSP");
    // Path to crypto materials.
	Path CRYPTO_PATH = Paths.get("../../test-network/organizations/peerOrganizations/org2.libBadge.com");
	// Path to user certificate.
	// Path to peer tls certificate.
	Path TLS_CERT_PATH = CRYPTO_PATH.resolve(Paths.get("peers/peer0.org2.libBadge.com/tls/ca.crt"));
	// Gateway peer end point.
	String PEER_ENDPOINT = "localhost:7051";
	String OVERRIDE_AUTH = "peer0.org2.libBadge.com";




    //미승인 상태 bogo 전체 조회
    //chaincode - GetUnapprovedBogos / evaluateTransaction(peer chaincode query)
    public ResponseEntity<Map<String, Object>> getUnapprovedBogoList(MemDto memDto) {
        Map<String, Object> result = new HashMap<>();

        List<Bogo> bogoList;

        Member member = memDto.toEntity();

        String email = member.getEmail();
        String name = member.getName();

        Lib lib = libRepository.findByMember_Email(email)
                .orElseThrow(() -> new RuntimeException("There is no lib information corresponding to email " + email));

        Long libNo = lib.getId();

        Path KEY_DIR_PATH = CRYPTO_PATH.resolve(Paths.get("users/user" + libNo + "@org1.libBadge.com/msp/keystore"));
        Path CERT_PATH = CRYPTO_PATH.resolve(Paths.get("users/user" + libNo + "@org1.libBadge.com/msp/signcerts/cert.pem"));

        try {
            Contract contract = HyperledgerFabricGateway.initializeContract(
                    MSP_ID, CRYPTO_PATH, TLS_CERT_PATH, KEY_DIR_PATH, CERT_PATH, PEER_ENDPOINT, OVERRIDE_AUTH
            );

            // Complete the code to query unapproved Bogos
            byte[] resultBytes = contract.evaluateTransaction("GetUnapprovedBogos");
            String resultString = new String(resultBytes, StandardCharsets.UTF_8);

            ObjectMapper objectMapper = new ObjectMapper();
            bogoList = Arrays.asList(objectMapper.readValue(resultString, Bogo[].class));

            List<BogoResponseDto> bogoResponseDtoList = new ArrayList<>();
            for (Bogo bogo : bogoList) {
                BogoResponseDto dto = new BogoResponseDto(bogo);
                bogoResponseDtoList.add(dto);
            }

            result.put("list", bogoResponseDtoList);

            return ResponseEntity.ok(result);
        } catch (GatewayException | JsonProcessingException e) {
            // Handle exceptions appropriately
            result.put("error", "An error occurred: " + e.getMessage());
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(result);
        }
    }

/* 
    //bogoNo 기준으로 상세 조회
    //chaincode - ReadBogo /evaluateTransaction(peer chaincode query)
    public BogoResponseDto getBogo(Integer bogoNo){
        Bogo bogo = contract.evaluateTransaction("ReadBogo", bogoNo);

        return new BogoResponseDto(bogo);
    }



    //특정 bogo에 대한 승인
    //chaincode - s / submitTransaction(peer chaincode invoke)



    //배지 발급
    //chaincode - IssueBadge / submitTransaction(peer chaincode invoke)
    //bogo에서 recipient (email가져와야 되는데 어떻게?)

    public type issueBadge(IssueRequestDto issueRequestDto){

        Badge badge = issueRequestDto.toEntity();

        String issuerId = badge.
        //String issuerName = 
        String recipient = i
    }
    
*/

}
