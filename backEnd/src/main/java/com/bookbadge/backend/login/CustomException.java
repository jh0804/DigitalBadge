package com.bookbadge.backend.login;

import lombok.Getter;

@Getter
public class CustomException extends RuntimeException{
    // @author 황승주
    private final CustomExceptionList exception;

    public CustomException(CustomExceptionList e) {
        super(e.getMessage());
        this.exception = e;
    }
}