package com.bookbadge.backend.badge;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Setter
@Getter
@AllArgsConstructor
public class BadgeResponseDto {
    	private Integer badgeNo;
        private String name; //배지 이름
        private String issueDate;

        //library 측 member
        private String issuerId;
        private String issuerName;

        //student 측 member
        private String recipient;

        private String level;
        private String description;
        
        /*Entity -> Dto*/
        public BadgeResponseDto(Badge badge) {
            this.badgeNo = badge.getBadgeNo();
            this.name = badge.getName();
            this.issueDate = badge.getIssueDate();
            /* 
            if (badge.getMember() != null) {
                this.issuerName = badge.getMember().getName();
                this.issuerId = badge.getMember().getEmail();
            }*/

            this.level = badge.getLevel();
            this.description = badge.getDescription();
        }

}
