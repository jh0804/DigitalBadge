package com.bookbadge.backend;

import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import com.bookbadge.backend.badge.Badge;
import com.bookbadge.backend.badge.BadgeResponseDto;
import com.bookbadge.backend.badge.BdgRequestDto;
import com.bookbadge.backend.bogo.Bogo;
import com.bookbadge.backend.bogo.BogoRequestDto;
import com.bookbadge.backend.bogo.BogoResponseDto;
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
public class StudentService {

    @Autowired
    private StuRepository stuRepository;

    private static boolean initialized = false; //IntiLedger를 위한
    
    String MSP_ID = System.getenv().getOrDefault("MSP_ID", "Org1MSP");
    // Path to crypto materials.
	Path CRYPTO_PATH = Paths.get("../../test-network/organizations/peerOrganizations/org1.libBadge.com");
	// Path to user certificate.
	// Path to peer tls certificate.
	Path TLS_CERT_PATH = CRYPTO_PATH.resolve(Paths.get("peers/peer0.org1.libBadge.com/tls/ca.crt"));
	// Gateway peer end point.
	String PEER_ENDPOINT = "localhost:7051";
	String OVERRIDE_AUTH = "peer0.org1.libBadge.com";

	//private final Gson gson = new GsonBuilder().setPrettyPrinting().create();


    //CreateBogo: bogo 생성
    ////chaincode - CreateBogo / submitTransaction(peer chaincode invoke)
    public BogoResponseDto createBogo(BogoRequestDto bogoRequestDto) {
        try {
            System.out.println("Entering createBogo");
    
            Bogo bogo = bogoRequestDto.toEntity();
    
            System.out.println("Bogo entity created");
    
            String title = bogo.getTitle();
            String author = bogo.getAuthor();
            String publisher = bogo.getPublisher();
            String report = bogo.getReport();
            String ownerId = bogo.getMember().getEmail();
            String ownerName = bogo.getMember().getName();
            //Boolean approved = bogo.getApproved();
    
            System.out.println("Title: " + title);
            System.out.println("Author: " + author);
            System.out.println("Publisher: " + publisher);
            System.out.println("Report: " + report);
            System.out.println("OwnerId: " + ownerId);
            System.out.println("OwnerName: " + ownerName);
            //System.out.println("Approved: " + approved);
    
            Stu stu = stuRepository.findByMember_Email(ownerId)
                    .orElseThrow(() -> new RuntimeException("There is no stu information corresponding to ownerId " + ownerId));
    
            Long stuNo = stu.getId();
    
            System.out.println("Stu information retrieved. StuNo: " + stuNo);
    
            Path KEY_DIR_PATH = CRYPTO_PATH.resolve(Paths.get("users/User" + stuNo + "@org1.libBadge.com/msp/keystore"));
            Path CERT_PATH = CRYPTO_PATH.resolve(Paths.get("users/User" + stuNo + "@org1.libBadge.com/msp/signcerts/cert.pem"));
    
            System.out.println("Key and Cert paths created");
    
            Contract contract = HyperledgerFabricGateway.initializeContract(
                    MSP_ID, CRYPTO_PATH, TLS_CERT_PATH, KEY_DIR_PATH, CERT_PATH, PEER_ENDPOINT, OVERRIDE_AUTH
            );
    
            System.out.println("Contract initialized");
    
            //InitLedger first run
            if(!initialized){
                byte[] initResult = contract.submitTransaction("InitLedger");
                initialized = true;
                System.out.println("InitLedger executed");
            }
    
            var result = contract.submitTransaction("CreateBogo", title, author, publisher, report, ownerId, ownerName);
    
            System.out.println("CreateBogo transaction submitted");
    
            String resultAsString = new String(result, StandardCharsets.UTF_8);
            int bogoNo = Integer.parseInt(resultAsString);
    
            System.out.println("BogoNo: " + bogoNo);
    
            BogoResponseDto bogoResponseDto = new BogoResponseDto(bogo);
    
            System.out.println("BogoResponseDto created");
    
            return bogoResponseDto;
        } catch (Exception e) {
            throw new RuntimeException("Failed to create Bogo", e);
        }
    }

    //Badge 목록 Recipient 기준 조회 
    ////chaincode - GetBadgesByRecipient / evaluateTransaction(peer chaincode query)
    public ResponseEntity<Map<String, Object>> getBadgeList(MemDto memDto) {
        Map<String, Object> result = new HashMap<>();
    
        List<Badge> badgeList;
    
        Member member = memDto.toEntity();
    
        String email = member.getEmail();
        String name = member.getName();
    
        Stu stu = stuRepository.findByMember_Email(email)
                .orElseThrow(() -> new RuntimeException(email + "에 해당하는 stu 정보가 없습니다."));
    
        Long stuNo = stu.getId();
    
        Path KEY_DIR_PATH = CRYPTO_PATH.resolve(Paths.get("users/user" + stuNo + "@org1.libBadge.com/msp/keystore"));
        Path CERT_PATH = CRYPTO_PATH.resolve(Paths.get("users/user" + stuNo + "@org1.libBadge.com/msp/signcerts/cert.pem"));
    
        try {
            Contract contract = HyperledgerFabricGateway.initializeContract(
                    MSP_ID, CRYPTO_PATH, TLS_CERT_PATH, KEY_DIR_PATH, CERT_PATH, PEER_ENDPOINT, OVERRIDE_AUTH
            );
    
            byte[] resultBytes = contract.evaluateTransaction("GetBadgesByRecipient", email);
            String resultString = new String(resultBytes, StandardCharsets.UTF_8);
    
            ObjectMapper objectMapper = new ObjectMapper();
            badgeList = Arrays.asList(objectMapper.readValue(resultString, Badge[].class));
    
            List<BadgeResponseDto> badgeResponseDtoList = new ArrayList<>();
            for (Badge badge : badgeList) {
                BadgeResponseDto dto = new BadgeResponseDto(badge);
                badgeResponseDtoList.add(dto);
            }
    
            result.put("list", badgeResponseDtoList);
    
            return ResponseEntity.ok(result);
        } catch (GatewayException | JsonProcessingException e) {
            result.put("error", "An error occurred: " + e.getMessage());
            return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(result);
        }
    }
    


    //Badge badgeBo 기준으로 상세 조회
    //chaincode - GetBadgeByBadgeNo / evaluateTransaction(peer chaincode query)
    public BadgeResponseDto getBadge(BdgRequestDto bdgRequestDto) {
        Badge badge = bdgRequestDto.toEntity();
    
        Member recipient = badge.getRecipient();
        int badgeNo = badge.getBadgeNo();
    
        Stu stu = stuRepository.findByMember_Email(recipient.getEmail())
        .orElseThrow(() -> new RuntimeException(recipient.getEmail() + "에 해당하는 stu 정보가 없습니다." ));
    
        Long stuNo = stu.getId();
    
        Path KEY_DIR_PATH = CRYPTO_PATH.resolve(Paths.get("users/user" + stuNo + "@org1.libBadge.com/msp/keystore"));
        Path CERT_PATH = CRYPTO_PATH.resolve(Paths.get("users/user" + stuNo + "@org1.libBadge.com/msp/signcerts/cert.pem"));
    
        try {
            Contract contract = HyperledgerFabricGateway.initializeContract(
                    MSP_ID, CRYPTO_PATH, TLS_CERT_PATH, KEY_DIR_PATH, CERT_PATH, PEER_ENDPOINT, OVERRIDE_AUTH
            );
    
            byte[] badgeNoBytes = String.valueOf(badgeNo).getBytes(StandardCharsets.UTF_8);

            byte[] result = contract.evaluateTransaction("GetBadgeByBadgeNo", badgeNoBytes);

            String resultAsString = new String(result, StandardCharsets.UTF_8);
            BadgeResponseDto badgeResponseDto = new Gson().fromJson(resultAsString, BadgeResponseDto.class);
    
            return badgeResponseDto;
        } catch (Exception e) {
            throw new RuntimeException("Failed to get badge", e);
        }
    }
}
