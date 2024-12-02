FROM gitpod/workspace-full
\

# alias all the things
RUN echo 'alias home="cd ${GITPOD_REPO_ROOT}"' | tee -a ~/.bashrc ~/.zshrc