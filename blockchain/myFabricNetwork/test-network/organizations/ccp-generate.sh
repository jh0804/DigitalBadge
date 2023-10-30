#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=library
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/library.libBadge.com/tlsca/tlsca.library.libBadge.com-cert.pem
CAPEM=organizations/peerOrganizations/library.libBadge.com/ca/ca.library.libBadge.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/library.libBadge.com/connection-library.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/library.libBadge.com/connection-library.yaml

ORG=student
P0PORT=9051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/student.libBadge.com/tlsca/tlsca.student.libBadge.com-cert.pem
CAPEM=organizations/peerOrganizations/student.libBadge.com/ca/ca.student.libBadge.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/student.libBadge.com/connection-student.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/student.libBadge.com/connection-student.yaml

