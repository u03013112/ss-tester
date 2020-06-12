FROM u03013112/ss-tester-base
COPY build/ss-tester /usr/local/bin/ss-tester
CMD [ "/usr/local/bin/ss-tester" ]