package com.bookbadge.backend.badge;

import lombok.NoArgsConstructor;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;

import com.bookbadge.backend.member.Member;
import com.fasterxml.jackson.annotation.JsonManagedReference;

@NoArgsConstructor
@AllArgsConstructor
@Builder 
@Data
public class Badge {

    //@Id
    private int badgeNo;    
     
	private String name;        
	private String issueDate;      
	private String issuerName;  
	private String image;          
	private String level;       
	private String description; 

    @ManyToOne(fetch = FetchType.LAZY, cascade = CascadeType.ALL, optional = false)
    @JsonManagedReference
    @JoinColumn(name = "issuerId", referencedColumnName = "email")
    private Member issuerId;

    @ManyToOne(fetch = FetchType.LAZY, cascade = CascadeType.ALL, optional = false)
    @JsonManagedReference
    @JoinColumn(name = "recipient", referencedColumnName = "email")
    private Member recipient;

}
