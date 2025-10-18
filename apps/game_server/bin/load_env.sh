#!/bin/bash

export APP_ENV=${APP_ENV:-development}

if [ -f ".env.${APP_ENV}" ]; then
  set -a
  source .env.${APP_ENV}
  set +a
fi
