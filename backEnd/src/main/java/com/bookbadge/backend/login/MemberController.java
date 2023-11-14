package com.bookbadge.backend.login;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/member") 
public class MemberController {


    @Autowired
    private MemberService memberService;

    @PutMapping("/update")
    public ResponseEntity<MemberResponseDto> updateMember(@RequestBody MemberRequestDto memberRequestDto) {

        String email = memberRequestDto.getEmail();
        MemberResponseDto updateMember = memberService.updateMember(email,memberRequestDto);
        return ResponseEntity.ok(updateMember);

    }
}
