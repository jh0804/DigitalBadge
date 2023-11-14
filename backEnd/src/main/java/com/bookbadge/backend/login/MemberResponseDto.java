package com.bookbadge.backend.login;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class MemberResponseDto {

    private String name;
    private String email;
    private String roleId;
    private Role role;

	/*Entity -> Dto*/
	public MemberResponseDto(Member member) {
		this.name = member.getName();
		this.email = member.getEmail();
		this.roleId = member.getRoleId();
		this.role = member.getRole();
	}
}