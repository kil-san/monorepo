FROM node:lts-alpine
ENV HOME=/srv/app
EXPOSE 4000
EXPOSE 8080
RUN apk --no-cache add openjdk11-jre bash
RUN yarn global add firebase-tools && yarn cache clean
RUN firebase setup:emulators:firestore
VOLUME $HOME/.cache
WORKDIR $HOME
CMD ["sh"]
