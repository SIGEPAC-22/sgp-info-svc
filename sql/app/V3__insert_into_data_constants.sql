
INSERT INTO ${flyway:database}.ste_state_data(ste_state_data_name) VALUES('ACTIVO');
INSERT INTO ${flyway:database}.ste_state_data(ste_state_data_name) VALUES('INACTIVO');


INSERT INTO ${flyway:database}.dct_document_type(dct_document_type_name) VALUES('N/A');
INSERT INTO ${flyway:database}.dct_document_type(dct_document_type_name) VALUES('DUI');
INSERT INTO ${flyway:database}.dct_document_type(dct_document_type_name) VALUES('NIT');
INSERT INTO ${flyway:database}.dct_document_type(dct_document_type_name) VALUES('PASAPORTE');


INSERT INTO ${flyway:database}.sex_patient(spt_gender_type) VALUES('MASCULINO');
INSERT INTO ${flyway:database}.sex_patient(spt_gender_type) VALUES('FEMENINA');


INSERT INTO ${flyway:database}.cty_country(cty_country_name) VALUES('EL SALVADOR');
INSERT INTO ${flyway:database}.cty_country(cty_country_name) VALUES('EXTRANJERO');

INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('AHUACHAPAN', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('CABAÑAS', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('CHALANTENANGO', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('CUSCATLAN', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('LA LIBERTAD', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('LA PAZ', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('LA UNION', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('MORAZAN', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('SAN MIGUEL', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('SAN SALVADOR', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('SAN VICENTE', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('SANTA ANA', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('SONSONATE', 1);
INSERT INTO ${flyway:database}.dpt_department(dpt_name_dapartment, dpt_id_country) VALUES('USULUTAN', 1);

INSERT INTO ${flyway:database}.spt_state_patient(spt_name_state_patient) VALUES('REVISION');
INSERT INTO ${flyway:database}.spt_state_patient(spt_name_state_patient) VALUES('ASINTOMATICO');
INSERT INTO ${flyway:database}.spt_state_patient(spt_name_state_patient) VALUES('ESTABLE');
INSERT INTO ${flyway:database}.spt_state_patient(spt_name_state_patient) VALUES('MODERADO');
INSERT INTO ${flyway:database}.spt_state_patient(spt_name_state_patient) VALUES('GRAVE');
INSERT INTO ${flyway:database}.spt_state_patient(spt_name_state_patient) VALUES('CRÍTICO');