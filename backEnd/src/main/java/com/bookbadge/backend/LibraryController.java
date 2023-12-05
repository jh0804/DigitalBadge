package com.bookbadge.backend;

import java.util.*;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.bookbadge.backend.badge.BadgeResponseDto;
import com.bookbadge.backend.bogo.BogoResponseDto;
import com.bookbadge.backend.member.MemDto;

@RestController
@RequestMapping("/admin")
public class LibraryController {
    
    private final LibraryService libraryService;

    @Autowired
    public LibraryController(LibraryService libraryService) {
        this.libraryService = libraryService;
    }
/* 
    //미승인 상태 bogo 전체 조회
    //chaincode - GetUnapprovedBogos
    @GetMapping("/bogo")
    
    public ResponseEntity<Map<String, Object>> getUnapprovedBogoList(@RequestBody MemDto memDto) {
 
        ResponseEntity<Map<String, Object>> response = libraryService.getUnapprovedBogoList(memDto);
        
        if (response == null || response.getBody() == null) {
            return ResponseEntity.ok(Collections.emptyList());
        }
        
        List<BogoResponseDto> bogoResponseDtoList = (List<BogoResponseDto>) response.getBody().get("list");
    
        return ResponseEntity.ok(bogoResponseDtoList);
    }


    //bogoNo 기준으로 상세 조회
    //chaincode - ReadBogo
    @GetMapping("/bogo/{bogoNo}")



    //특정 bogo에 대한 승인
    //chaincode - ApproveBogo
    @PostMapping("/bogo/{bogoNo}")



    //배지 발급
    //chaincode - IssueBadge
    @PostMapping("/badge")
*/

}
