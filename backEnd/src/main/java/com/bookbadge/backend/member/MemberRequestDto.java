package com.bookbadge.backend.member;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class MemberRequestDto {

    private String name;
    private String email;
    private String roleId;
    private Role role;

    /* Dto -> Entity */
	public Member toEntity() {
        Member member = Member.builder()
                    .name(name)
                    .email(email)
                    .roleId(roleId)
                    .role(role)
                    .build();
        return member;
	}
}