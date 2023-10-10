package com.bookbadge.backend.user;

import javax.persistence.*;
import org.hibernate.annotations.DynamicUpdate;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@DynamicUpdate //update 할때 실제 값이 변경됨 컬럼으로만 update 쿼리를 만듦
@Entity 
@Getter 
@Table(name = "user") 
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "uId")
    private Long id;

    @Column(name = "name", nullable = false)
    private String name;

    @Column(name = "email", nullable = false)
    private String email;

    @Column(name = "picture")
    private String picture;
    
    @Enumerated(EnumType.STRING)
    @Column(name = "role", nullable = false)
    private Role role; //유저 권한 (ADMIN, USER)

    @Column(name = "univ")
    private String univ;

    @Column(name = "major")
    private String major;
    
    @Column(name = "grade")
    private String grade;

    @Column(name = "stuId")
    private String stuId;

    @Column(name = "team")
    private String team;


    @Builder
    public User(String name, String email, String picture, Role role) {
        this.name = name;
        this.email = email;
        this.picture = picture;
        this.role = role;
    }

    public User update(String name, String picture, String univ, String major,String grade, String stuId,String team) {
        this.name = name;
        this.picture = picture;
        this.univ = univ;
        this.major = major;
        this.grade = grade;
        this.stuId = stuId;
        this.team = team;

        return this;
    }

    public String getRoleKey() {
        return this.role.getKey();
    }
}
