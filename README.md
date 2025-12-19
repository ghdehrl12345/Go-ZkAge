# Go-ZkAge


1. Circuit Definition : Define 함수에서 수학적 제약 조건(A <= B)을 설계
2. compilation : 회로를 R1CS 제약 시스템으로 변환
3. Setup : 증명키(pk)와 검증키(vk)를 생성
4. Witness Generation : 비밀 데이터와 공개 데이터를 할당
5. Proving : groth16.Prove로 증명서(proof)를 생성
6. Verification : groth16.Verify로 증명서의 유효성을 검증

설명 : https://www.notion.so/ZAGE-Go-2cbdae8d548580f79210f06ce2f1a5a9?source=copy_link