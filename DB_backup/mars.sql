PGDMP  (    +        	        |            mars    16.1    16.1     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    25012    mars    DATABASE     x   CREATE DATABASE mars WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';
    DROP DATABASE mars;
                postgres    false            �            1259    25013    report    TABLE     �   CREATE TABLE public.report (
    id integer NOT NULL,
    report_text character varying,
    file_path path,
    name character varying NOT NULL
);
    DROP TABLE public.report;
       public         heap    postgres    false            P           2606    25019    report report_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.report
    ADD CONSTRAINT report_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.report DROP CONSTRAINT report_pkey;
       public            postgres    false    215           