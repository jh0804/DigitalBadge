package com.bookbadge.backend.member;

import lombok.*;
import javax.persistence.*;

import org.hibernate.annotations.DynamicUpdate;

@NoArgsConstructor(access = AccessLevel.PROTECTED) 
@AllArgsConstructor
@DynamicUpdate 
@Data
@Entity
@Table(name = "lib") 
public class Lib {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "libNo")
    private Long id;

    @OneToOne
    @JoinColumn(name = "email", referencedColumnName = "email")
    private Member member;

    public Lib(Member member) {
        this.member = member;
    }

}