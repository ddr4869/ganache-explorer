# PostgreSQL Dockerfile

# 베이스 이미지를 정의합니다. 여기서는 공식 PostgreSQL 이미지를 사용합니다.
FROM postgres:latest

# 환경 변수를 설정하여 PostgreSQL 데이터베이스를 생성하고 사용자를 만듭니다.
ENV POSTGRES_DB=test_apps
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=1821

# 컨테이너 내부의 5432 포트를 호스트의 5432 포트에 매핑합니다.
EXPOSE 5432