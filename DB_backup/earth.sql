PGDMP                      |            earth    16.1    16.1     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    25020    earth    DATABASE     y   CREATE DATABASE earth WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE earth;
                postgres    false            �            1259    25021    reports    TABLE     �   CREATE TABLE public.reports (
    id integer NOT NULL,
    file_name character varying,
    report_text character varying NOT NULL,
    reporter_name character varying NOT NULL
);
    DROP TABLE public.reports;
       public         heap    postgres    false            �           0    0    TABLE reports    ACL     G   GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE public.reports TO tm_admin;
          public          postgres    false    215            P           2606    25027    reports reports_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.reports
    ADD CONSTRAINT reports_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.reports DROP CONSTRAINT reports_pkey;
       public            postgres    false    215           