FROM alpine as CERTS
RUN apk add ca-certificates && \
    update-ca-certificates
FROM scratch
ENV DISCORD_URL=""
ENV DISCORD_AVATAR=""
ENV DISCORD_USERNAME=""
COPY --from=CERTS /etc/ssl/certs /etc/ssl/certs
COPY bin/discord-notif /usr/bin/discord-notif
USER 6000
ENTRYPOINT ["discord-notif"]
