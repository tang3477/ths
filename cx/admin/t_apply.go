--
-- PostgreSQL database dump
--

-- Dumped from database version 9.1.24
-- Dumped by pg_dump version 9.1.24
-- Started on 2017-06-02 11:29:45

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 169 (class 1259 OID 17703)
-- Dependencies: 1763 6
-- Name: t_apply; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE t_apply (
    id integer NOT NULL,
    lx character varying(50) NOT NULL,
    nam character varying(50) NOT NULL,
    dire character varying(50) NOT NULL,
    ppe character varying(50) NOT NULL,
    cntent character varying(500) NOT NULL,
    tim double precision NOT NULL,
    account integer NOT NULL,
    adopt integer DEFAULT 2 NOT NULL,
    dizhi character varying(100)
);


ALTER TABLE public.t_apply OWNER TO postgres;

--
-- TOC entry 168 (class 1259 OID 17701)
-- Dependencies: 6 169
-- Name: t_apply_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE t_apply_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_apply_id_seq OWNER TO postgres;

--
-- TOC entry 1873 (class 0 OID 0)
-- Dependencies: 168
-- Name: t_apply_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE t_apply_id_seq OWNED BY t_apply.id;


--
-- TOC entry 1762 (class 2604 OID 17706)
-- Dependencies: 168 169 169
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY t_apply ALTER COLUMN id SET DEFAULT nextval('t_apply_id_seq'::regclass);


--
-- TOC entry 1868 (class 0 OID 17703)
-- Dependencies: 169 1869
-- Data for Name: t_apply; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY t_apply (id, lx, nam, dire, ppe, cntent, tim, account, adopt, dizhi) FROM stdin;
13	我反倒没人来看	双方都没	外来人	为了人民	网络名人	1495911818	123456	1	\N
11	我发现	玩饥饿	有着	吴若甫	没胃口	1495911402	123456	1	\N
12	温柔的	内燃机	为荣	万通	我们相聚	1495911596	123456	1	qqq.rar
14	日常工具	万年历	程序	王五	一个小小的日常工具	1495931895	123	2	\N
15	游戏	贪吃蛇	益智	王五	一款小游戏	1495932734	123	1	贪吃蛇项目.rar
\.


--
-- TOC entry 1874 (class 0 OID 0)
-- Dependencies: 168
-- Name: t_apply_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('t_apply_id_seq', 15, true);


--
-- TOC entry 1765 (class 2606 OID 17711)
-- Dependencies: 169 169 1870
-- Name: t_apply_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY t_apply
    ADD CONSTRAINT t_apply_pkey PRIMARY KEY (id);


-- Completed on 2017-06-02 11:29:45

--
-- PostgreSQL database dump complete
--

