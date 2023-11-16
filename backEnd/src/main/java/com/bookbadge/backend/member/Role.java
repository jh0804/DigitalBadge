package com.bookbadge.backend.member;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum Role {

    UNKNOWN("ROLE_UNKNOWN","유저"),
    STUDENT("ROLE_STUDENT","학생"),
    LIBRARY("ROLE_LIBRARY","도서관 관리자");

    private final String key;
    private final String title;
}
