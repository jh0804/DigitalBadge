package com.bookbadge.backend.member;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

public interface MemberRepository extends JpaRepository<Member, Long> {
    Optional<Member> findByEmailAndProvider(String email, String provider);
    Optional<Member> findByEmail(String email);
    Optional<Member> findByEmail(Member email);
}
