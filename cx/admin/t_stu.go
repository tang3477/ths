--
-- PostgreSQL database dump
--

-- Dumped from database version 9.1.24
-- Dumped by pg_dump version 9.1.24
-- Started on 2017-06-02 11:30:08

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 163 (class 1259 OID 17679)
-- Dependencies: 6
-- Name: t_stu; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE t_stu (
    id integer NOT NULL,
    account integer NOT NULL,
    pwd integer NOT NULL,
    tel integer NOT NULL,
    e_mail character varying(20) NOT NULL
);


ALTER TABLE public.t_stu OWNER TO postgres;

--
-- TOC entry 162 (class 1259 OID 17677)
-- Dependencies: 163 6
-- Name: t_stu_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE t_stu_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_stu_id_seq OWNER TO postgres;

--
-- TOC entry 1872 (class 0 OID 0)
-- Dependencies: 162
-- Name: t_stu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE t_stu_id_seq OWNED BY t_stu.id;


--
-- TOC entry 1762 (class 2604 OID 17682)
-- Dependencies: 162 163 163
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY t_stu ALTER COLUMN id SET DEFAULT nextval('t_stu_id_seq'::regclass);


--
-- TOC entry 1867 (class 0 OID 17679)
-- Dependencies: 163 1868
-- Data for Name: t_stu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY t_stu (id, account, pwd, tel, e_mail) FROM stdin;
11	123456	123456	123456	1234@.com
12	123	123	123	123
\.


--
-- TOC entry 1873 (class 0 OID 0)
-- Dependencies: 162
-- Name: t_stu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('t_stu_id_seq', 12, true);


--
-- TOC entry 1764 (class 2606 OID 17684)
-- Dependencies: 163 163 1869
-- Name: t_stu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY t_stu
    ADD CONSTRAINT t_stu_pkey PRIMARY KEY (id);


-- Completed on 2017-06-02 11:30:09

--
-- PostgreSQL database dump complete
--

