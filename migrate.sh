#!/bin/bash
rm nihao.db
touch nihao.db
chmod 777 nihao.db
docker compose down && docker compose up -d