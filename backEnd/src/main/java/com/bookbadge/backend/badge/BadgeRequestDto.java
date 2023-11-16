package com.bookbadge.backend.badge;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class BadgeRequestDto {
    private String recipient; //학생 email
    private String issuerId; //
}
