PGDMP                         u            csngdzb    9.1.24    9.1.24 	    N           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            O           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            �            1259    17687    t_adm    TABLE     �   CREATE TABLE t_adm (
    id integer NOT NULL,
    account integer NOT NULL,
    pwd integer NOT NULL,
    tel integer NOT NULL,
    e_mail character varying(20) NOT NULL,
    code character varying(100) NOT NULL
);
    DROP TABLE public.t_adm;
       public         postgres    false    6            �            1259    17685    t_adm_id_seq    SEQUENCE     n   CREATE SEQUENCE t_adm_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.t_adm_id_seq;
       public       postgres    false    165    6            P           0    0    t_adm_id_seq    SEQUENCE OWNED BY     /   ALTER SEQUENCE t_adm_id_seq OWNED BY t_adm.id;
            public       postgres    false    164            �           2604    17690    id    DEFAULT     V   ALTER TABLE ONLY t_adm ALTER COLUMN id SET DEFAULT nextval('t_adm_id_seq'::regclass);
 7   ALTER TABLE public.t_adm ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    165    164    165            K          0    17687    t_adm 
   TABLE DATA               =   COPY t_adm (id, account, pwd, tel, e_mail, code) FROM stdin;
    public       postgres    false    165    1868   �       Q           0    0    t_adm_id_seq    SEQUENCE SET     3   SELECT pg_catalog.setval('t_adm_id_seq', 3, true);
            public       postgres    false    164            �           2606    17692 
   t_adm_pkey 
   CONSTRAINT     G   ALTER TABLE ONLY t_adm
    ADD CONSTRAINT t_adm_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.t_adm DROP CONSTRAINT t_adm_pkey;
       public         postgres    false    165    165    1869            K   /   x�3�44172��SF�f�&P�A/9?�����������А+F��� �	�     