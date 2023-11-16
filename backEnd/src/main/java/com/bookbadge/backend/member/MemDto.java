package com.bookbadge.backend.member;

public class MemDto {
    
    private String email;
    private String name;

    
    /* Dto -> Entity */
	public Member toEntity() {
        Member member = Member.builder()
                    .name(name)
                    .email(email)
                    .build();
        return member;
	}
    
}