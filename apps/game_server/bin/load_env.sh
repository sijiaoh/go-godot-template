#!/bin/bash

export APP_ENV=${APP_ENV:-development}

if [ -f ".env.${APP_ENV}" ]; then
  source .env.${APP_ENV}
fi
