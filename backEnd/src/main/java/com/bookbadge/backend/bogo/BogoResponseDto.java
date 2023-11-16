package com.bookbadge.backend.bogo;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
@AllArgsConstructor
public class BogoResponseDto {

	private Integer bogoNo;
	private String title;
	private String author;
	private String publisher;
	private String report;
	private Boolean approved;

	//member email&name 
	private String email;
	private String name;

	/*Entity -> Dto*/
	public BogoResponseDto(Bogo bogo) {
		this.bogoNo = bogo.getBogoNo();
		this.title = bogo.getTitle();
		this.author = bogo.getAuthor();
		this.publisher = bogo.getPublisher();
		this.report = bogo.getReport();
		if (bogo.getMember() != null) {
			this.name = bogo.getMember().getName();
			this.email = bogo.getMember().getEmail();
		}
		this.approved = bogo.getApproved();
	}

	@Builder 
	public BogoResponseDto(Integer bogoNo, String title, String author, String publisher, String report, String email, String name, Boolean approved){
		this.bogoNo = bogoNo;
		this.title = title;
		this.author = author;
		this.publisher = publisher;
		this.report = report;
		this.email = email;
		this.approved = approved;
	}

}