PGDMP                         u            csngdzb    9.1.24    9.1.24 	    N           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            O           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            �            1259    17695    t_aid    TABLE     Q   CREATE TABLE t_aid (
    id integer NOT NULL,
    code character varying(100)
);
    DROP TABLE public.t_aid;
       public         postgres    false    6            �            1259    17693    t_aid_id_seq    SEQUENCE     n   CREATE SEQUENCE t_aid_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.t_aid_id_seq;
       public       postgres    false    6    167            P           0    0    t_aid_id_seq    SEQUENCE OWNED BY     /   ALTER SEQUENCE t_aid_id_seq OWNED BY t_aid.id;
            public       postgres    false    166            �           2604    17698    id    DEFAULT     V   ALTER TABLE ONLY t_aid ALTER COLUMN id SET DEFAULT nextval('t_aid_id_seq'::regclass);
 7   ALTER TABLE public.t_aid ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    166    167    167            K          0    17695    t_aid 
   TABLE DATA               "   COPY t_aid (id, code) FROM stdin;
    public       postgres    false    167    1868   �       Q           0    0    t_aid_id_seq    SEQUENCE SET     4   SELECT pg_catalog.setval('t_aid_id_seq', 1, false);
            public       postgres    false    166            �           2606    17700 
   t_aid_pkey 
   CONSTRAINT     G   ALTER TABLE ONLY t_aid
    ADD CONSTRAINT t_aid_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.t_aid DROP CONSTRAINT t_aid_pkey;
       public         postgres    false    167    167    1869            K   /   x�M�9 0�:&c�C�8�j���VeGH"mIي�5cֶ�|��1     