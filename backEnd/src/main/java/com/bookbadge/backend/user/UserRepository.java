package com.bookbadge.backend.user;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByEmail(String email); //email 기준으로 이미 존재하는 회원인지 확인
}
