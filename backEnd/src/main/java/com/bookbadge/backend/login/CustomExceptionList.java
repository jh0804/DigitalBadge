package com.bookbadge.backend.login;

import lombok.Getter;
import lombok.ToString;
import org.springframework.http.HttpStatus;

@Getter
@ToString
public enum CustomExceptionList {

    RUNTIME_EXCEPTION(HttpStatus.BAD_REQUEST, "E001", "잘못된 요청입니다."),
    INTERNAL_SERVER_ERROR(HttpStatus.INTERNAL_SERVER_ERROR, "E002", "서버 오류 입니다."),
    ACCESS_TOKEN_ERROR(HttpStatus.UNAUTHORIZED, "E003", "엑세스 토큰 오류입니다."),
    REFRESH_TOKEN_ERROR(HttpStatus.UNAUTHORIZED, "E004", "리프레쉬 토큰 오류입니다."),
    MEMBER_NOT_FOUND_ERROR(HttpStatus.NOT_FOUND, "E005", "존재하지 않는 회원입니다."),
    JOIN_INFO_NOT_EXIST(HttpStatus.NOT_FOUND, "E008", "가입정보가 유효하지 않습니다."),
    NO_AUTHENTICATION_ERROR(HttpStatus.FORBIDDEN, "E009", "접근 권한이 없습니다.");

    private final HttpStatus status;
    private final String code;
    private String message;

    CustomExceptionList(HttpStatus status, String code) {
        this.status = status;
        this.code = code;
    }

    CustomExceptionList(HttpStatus status, String code, String message) {
        this.status = status;
        this.code = code;
        this.message = message;
    }
}
