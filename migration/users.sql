-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";

CREATE TABLE "public"."users" (
  "id" SERIAL,
  "name" VARCHAR(70),
  "email" VARCHAR(70) UNIQUE,
  "created_at" TIMESTAMP,
  "updated_at" TIMESTAMP
);

ALTER TABLE "public"."users" OWNER TO "atzgywhtayemui";

-- ----------------------------
-- Primary Key for table orders
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "user_id_pkey" PRIMARY KEY ("id");
