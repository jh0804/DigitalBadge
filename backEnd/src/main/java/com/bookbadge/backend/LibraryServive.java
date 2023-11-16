package com.bookbadge.backend;

import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import com.bookbadge.backend.badge.Badge;
import com.bookbadge.backend.bogo.Bogo;
import com.bookbadge.backend.bogo.BogoResponseDto;
import com.bookbadge.backend.member.Member;

@Service
public class LibraryServive {

    private String MSP_ID = System.getenv().getOrDefault("MSP_ID", "Org2MSP");
	private static final String CHANNEL_NAME = System.getenv().getOrDefault("CHANNEL_NAME", "mychannel");
	private static final String CHAINCODE_NAME = System.getenv().getOrDefault("CHAINCODE_NAME", "basic");
    // Path to crypto materials.
	private static final Path CRYPTO_PATH = Paths.get("../../test-network/organizations/peerOrganizations/org2.libBadge.com");
	// Path to user certificate.
	// Path to peer tls certificate.
	private static final Path TLS_CERT_PATH = CRYPTO_PATH.resolve(Paths.get("peers/peer0.org2.libBadge.com/tls/ca.crt"));
	// Gateway peer end point.
	private static final String PEER_ENDPOINT = "localhost:9051";
	private static final String OVERRIDE_AUTH = "peer0.org2.libBadge.com";



    /* 

    //미승인 상태 bogo 전체 조회
    //chaincode - GetUnapprovedBogos / evaluateTransaction(peer chaincode query)

    public ResponseEntity<Map<String, Object>> getUnapprovedBogoList(){
        Map<String, Object> result = new HashMap<>();

        List<Bogo> bogoList;

        bogoList = contract.evaluateTransaction("GetUnapprovedBogos");
        List<BogoResponseDto> bogoResponseDtoList = new ArrayList<>();
                for (Bogo bogo : bogoList) {
                    BogoResponseDto dto = new BogoResponseDto(bogo);
                    bogoResponseDtoList.add(dto);
                }

        result.put("list", bogoResponseDtoList);

        return ResponseEntity.ok(result);
    }

    //bogoNo 기준으로 상세 조회
    //chaincode - ReadBogo /evaluateTransaction(peer chaincode query)
    public BogoResponseDto getBogo(Integer bogoNo){
        Bogo bogo = contract.evaluateTransaction("ReadBogo", bogoNo);

        return new BogoResponseDto(bogo);
    }



    //특정 bogo에 대한 승인
    //chaincode - ApproveBogo / submitTransaction(peer chaincode invoke)



    //배지 발급
    //chaincode - IssueBadge / submitTransaction(peer chaincode invoke)
    IssueBadge(ctx contractapi.TransactionContextInterface, recipient string, issuerId string)

    //bogo에서 recipient (email가져와야 되는데 어떻게?)

    public type issueBadge(IssueRequestDto issueRequestDto){

        Badge badge = issueRequestDto.toEntity();

        String issuerId = badge.
        String issuerName = 
        String recipient = i
    }
    

*/

}
