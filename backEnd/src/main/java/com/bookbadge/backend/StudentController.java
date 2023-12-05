package com.bookbadge.backend;

import java.util.*;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.bookbadge.backend.badge.BadgeResponseDto;
import com.bookbadge.backend.badge.BdgRequestDto;
import com.bookbadge.backend.bogo.BogoRequestDto;
import com.bookbadge.backend.bogo.BogoResponseDto;
import com.bookbadge.backend.member.MemDto;

@RestController
@RequestMapping("/student")
public class StudentController {
    
    @Autowired
    private StudentService studentService; 

    //bogo 등록
    //chaincode - CreateBogo
    @PostMapping("/bogo")
    public ResponseEntity<BogoResponseDto> createBogo(@RequestBody BogoRequestDto bogoRequestDto) {
        BogoResponseDto bogoResponseDto = studentService.createBogo(bogoRequestDto);
        return ResponseEntity.ok(bogoResponseDto);

    }
    

    //Recipient 기준으로 badge 목록 조회
    //chaincode - GetBadgesByRecipient
    @GetMapping("/badge")
    public ResponseEntity<List<BadgeResponseDto>> getBadgesByRecipient(@RequestBody MemDto memDto) {

        ResponseEntity<Map<String, Object>> response = studentService.getBadgeList(memDto);

        if (response == null || response.getBody() == null) {
            return ResponseEntity.ok(Collections.emptyList());
        }
        
        List<BadgeResponseDto> badgeResponseDtoList = (List<BadgeResponseDto>) response.getBody().get("list");
    
        return ResponseEntity.ok(badgeResponseDtoList);

    }

    //badgeNo 기준으로 badge 상세 조회
    //chaincode - GetBadgeByBadgeNo
    @GetMapping("/bogo/{bogoNo}")
    public ResponseEntity<BadgeResponseDto> GetBadgeByBadgeNo(@RequestBody BdgRequestDto bdgRequestDto) {
        BadgeResponseDto badgeResponseDto = studentService.getBadge(bdgRequestDto);
        return ResponseEntity.ok(badgeResponseDto);
    }



}
