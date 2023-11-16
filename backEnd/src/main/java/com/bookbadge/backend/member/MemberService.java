package com.bookbadge.backend.member;

import javax.transaction.Transactional;
import org.springframework.stereotype.Service;

import com.bookbadge.backend.bogo.Bogo;

import lombok.RequiredArgsConstructor;


@Service
@RequiredArgsConstructor
@Transactional
public class MemberService {

    private final MemberRepository memberRepository;
    private final StuRepository stuRepository;
    private final LibRepository libRepository;

    public MemberResponseDto updateMember(MemberRequestDto memberRequestDto) {
        Member member = memberRepository.findByEmail(memberRequestDto.getEmail())
                .orElseThrow(() -> new IllegalArgumentException("해당 유저는 존재하지 않습니다: "));

        member.setRoleId(memberRequestDto.getRoleId());
        member.setRole(memberRequestDto.getRole());
    
        Member updateMember = memberRepository.save(member);
    
        if (Role.STUDENT.equals(memberRequestDto.getRole())) {
            boolean emailExistsInStu = stuRepository.existsByMember_Email(memberRequestDto.getEmail());

            if (!emailExistsInStu) {
                Stu stu = new Stu(updateMember);
                stuRepository.save(stu);
            }
        } else if (Role.LIBRARY.equals(memberRequestDto.getRole())) {
            boolean emailExistsInLib = libRepository.existsByMember_Email(memberRequestDto.getEmail());

            if (!emailExistsInLib) {
                Lib lib = new Lib(updateMember);
                libRepository.save(lib);
            }
        }
    
        return new MemberResponseDto(updateMember);
    }
    
}

