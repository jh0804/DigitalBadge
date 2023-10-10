package com.bookbadge.backend.user;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public enum Role {

   ADMIN("ROLE_ADMIN", "관리자"),
   USER("ROLE_USER", "학생");

   private final String key;
   private final String title;

}
