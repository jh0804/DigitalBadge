package com.bookbadge.backend.login;

import java.time.LocalDateTime;

import javax.transaction.Transactional;

import org.springframework.stereotype.Service;

import lombok.RequiredArgsConstructor;


@Service
@RequiredArgsConstructor
@Transactional
public class MemberService {

    private final MemberRepository memberRepository;


    public MemberResponseDto updateMember(String email, MemberRequestDto memberRequestDto) {
        Member member = memberRepository.findByEmail(email)
                .orElseThrow(() -> new IllegalArgumentException("다음의 이메일에 해당하는 유저가 존재하지 않습니다: " + email));

        member.setRoleId(memberRequestDto.getRoleId());
        member.setRole(memberRequestDto.getRole());
		Member updateMember = memberRepository.save(member);
		
        return new MemberResponseDto(updateMember);
    }
}
