package com.bookbadge.backend.member;

import lombok.*;
import javax.persistence.*;

import org.hibernate.annotations.DynamicUpdate;

@NoArgsConstructor(access = AccessLevel.PROTECTED)
@AllArgsConstructor
@DynamicUpdate
@Entity
@Data
@Table(name = "stu")
public class Stu {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "stuNo")
    private Long id;

    @OneToOne
    @JoinColumn(name = "email", referencedColumnName = "email")
    private Member member;

    public Stu(Member member) {
        this.member = member;
    }

}