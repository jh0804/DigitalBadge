package com.bookbadge.backend.bogo;


import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;

import org.hibernate.annotations.DynamicInsert;
import org.hibernate.annotations.DynamicUpdate;
import org.springframework.data.annotation.Id;

import com.bookbadge.backend.member.Member;
import com.fasterxml.jackson.annotation.JsonManagedReference;

import lombok.*;

@NoArgsConstructor
@AllArgsConstructor
@Builder 
@Data
public class Bogo {

    //@Id
    private int bogoNo;

    private String title;
    private String author;
    private String publisher;
    private String report;

    private Boolean approved;

    //member email&name
    @ManyToOne(fetch = FetchType.LAZY)
    @JsonManagedReference
    @JoinColumn(name="email") 
    private Member member;


}
