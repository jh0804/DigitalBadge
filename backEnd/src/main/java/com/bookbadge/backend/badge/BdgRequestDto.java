package com.bookbadge.backend.badge;


import com.bookbadge.backend.bogo.Bogo;
import com.bookbadge.backend.member.Member;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
//학생 배지 상세 조회용(email+badgeNo 필요)
public class BdgRequestDto {

    private String recipient; //학생 email (recipient)
    private int badgeNo; //badgeNo


        /* Dto -> Entity */ 
	public Badge toEntity() {
                Member member = Member.builder()
                        .email(recipient)
                        .build();
        
                return Badge.builder()
                        .badgeNo(badgeNo)
                        .recipient(member)
                        .build();
            }
}
