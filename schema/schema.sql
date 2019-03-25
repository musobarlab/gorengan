CREATE TABLE "PRODUCTS" (
    "ID" character varying(10) UNIQUE PRIMARY KEY NOT NULL, 
    "NAME" character varying(50) NOT NULL, 
    "QUANTITY" integer NOT NULL NOT NULL,
    "CATEGORY_ID" character varying(10) NOT NULL,
    "CREATOR_ID" character varying(20),
    "CREATOR_IP" character varying(20),
    "EDITOR_ID" character varying(20),
    "EDITOR_IP" character varying(20),
    "IS_DELETED" boolean DEFAULT false,
    "CREATED" timestamp with time zone DEFAULT now() NOT NULL, 
    "LAST_MODIFIED" timestamp with time zone,
    "DELETED" timestamp with time zone
);

CREATE TABLE "PRODUCT_CATEGORIES" (
    "ID" character varying(10) UNIQUE PRIMARY KEY NOT NULL, 
    "NAME" character varying(50) NOT NULL,
    "CREATOR_ID" character varying(20),
    "CREATOR_IP" character varying(20),
    "EDITOR_ID" character varying(20),
    "EDITOR_IP" character varying(20),
    "IS_DELETED" boolean DEFAULT false,
    "CREATED" timestamp with time zone DEFAULT now() NOT NULL, 
    "LAST_MODIFIED" timestamp with time zone,
    "DELETED" timestamp with time zone
);