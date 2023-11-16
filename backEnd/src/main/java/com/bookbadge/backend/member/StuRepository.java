package com.bookbadge.backend.member;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;

public interface StuRepository extends JpaRepository<Stu, Long>{

    Optional<Stu> findByMember_Email(String email);
    boolean existsByMember_Email(String email);

}

