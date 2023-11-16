package com.bookbadge.backend.bogo;



import com.bookbadge.backend.member.Member;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class BogoRequestDto {

	private String title;
	private String author;
	private String publisher;
	private String report;
        
        //memebr
        private String name;
	private String email;


        /* Dto -> Entity */ 
	public Bogo toEntity() {
                Member member = Member.builder()
                        .name(name)
                        .email(email)
                        .build();
                Bogo bogo = Bogo.builder()
                        .title(title)
                        .author(author)
                        .publisher(publisher)
                        .report(report)
                        .member(member)
                        .build();
                return bogo;
	}
}