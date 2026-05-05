FROM quay.io/keycloak/keycloak:26.6.0

WORKDIR /opt/keycloak

COPY testdata /opt/keycloak/data/import/

ENV KEYCLOAK_ADMIN=admin
ENV KEYCLOAK_ADMIN_PASSWORD=secret

ENV KC_HOSTNAME=localhost
ENV KC_HEALTH_ENABLED=true
ENV KC_FEATURES=account-api,authorization,client-policies,impersonation,docker,scripts,admin-fine-grained-authz,organization
RUN /opt/keycloak/bin/kc.sh build

ENTRYPOINT ["/opt/keycloak/bin/kc.sh"]
CMD ["start-dev", "--import-realm"]