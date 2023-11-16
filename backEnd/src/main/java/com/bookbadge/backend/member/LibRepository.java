package com.bookbadge.backend.member;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

public interface LibRepository extends JpaRepository<Lib, Long>{

    Optional<Lib> findByMember_Email(String email);
    boolean existsByMember_Email(String email);
}