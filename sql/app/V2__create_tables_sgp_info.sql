
create table ${flyway:database}.ste_state_data
(
    ste_id_state_data int auto_increment primary key not null,
    ste_state_data_name varchar(45) unique not null,
    ste_state_data_creation_date timestamp null,
    ste_state_data_created_by varchar(45) null,
    ste_state_data_modification_date timestamp null,
    ste_state_data_modified_by varchar(45) null
);

create trigger ${flyway:database}.ste_state_insert_aud
    BEFORE INSERT ON ${flyway:database}.ste_state_data
    FOR EACH ROW
    set NEW.ste_state_data_created_by=USER(),
	NEW.ste_state_data_creation_date=now();

create trigger ${flyway:database}.ste_state_update_aud
    BEFORE UPDATE ON ${flyway:database}.ste_state_data
    FOR EACH ROW
    set NEW.ste_state_data_modification_date=now(),
        NEW.ste_state_data_modified_by=USER();


create table ${flyway:database}.dct_document_type
(
    dct_id_document_type int auto_increment primary key,
    dct_document_type_name varchar(45) not null,
    dct_id_state_data int null default 1,
    dct_creation_date timestamp null,
    dct_created_by varchar(45) null,
    dct_modification_date timestamp null,
    dct_modified_by varchar(45) null,
    foreign key (dct_id_state_data) references ste_state_data(ste_id_state_data)
);

create trigger ${flyway:database}.dct_document_type_insert_aud
    BEFORE INSERT ON ${flyway:database}.dct_document_type
    FOR EACH ROW
    set NEW.dct_created_by=USER(),
	NEW.dct_creation_date=now();

create trigger ${flyway:database}.dct_document_type_update_aud
    BEFORE UPDATE ON ${flyway:database}.dct_document_type
    FOR EACH ROW
    set NEW.dct_modification_date=now(),
        NEW.dct_modified_by=USER();


create table ${flyway:database}.sex_patient
(
    spt_id_sex int auto_increment primary key not null,
    spt_gender_type varchar(45) not null,
    spt_id_state_data int null default 1,
    spt_creation_date timestamp null,
    spt_created_by varchar(45) null,
    spt_modification_date timestamp null,
    spt_modified_by varchar(45) null,
    foreign key (spt_id_state_data) references ste_state_data(ste_id_state_data)
);

create trigger ${flyway:database}.sex_patient_insert_aud
    BEFORE INSERT ON ${flyway:database}.sex_patient
    FOR EACH ROW
    set NEW.spt_created_by=USER(),
		NEW.spt_creation_date=now();

create trigger ${flyway:database}.sex_patient_update_aud
    BEFORE UPDATE ON ${flyway:database}.sex_patient
    FOR EACH ROW
    set NEW.spt_modification_date=now(),
        NEW.spt_modified_by=USER();



create table ${flyway:database}.stm_symptom
(
    stm_id_sympton int auto_increment primary key,
    stm_name_symptons varchar(45) unique not null,
    stm_sympton_description varchar(100) null,
    stm_id_state_data int null default 1,
    stm_creation_date timestamp null,
    stm_created_by varchar(45) null,
    stm_modification_date timestamp null,
    stm_modified_by varchar(45) null,
    foreign key (stm_id_state_data) references ste_state_data(ste_id_state_data)
);

create trigger ${flyway:database}.stm_symptom_insert_aud
    BEFORE INSERT ON ${flyway:database}.stm_symptom
    FOR EACH ROW
    set NEW.stm_created_by=USER(),
		NEW.stm_creation_date=now();

create trigger ${flyway:database}.stm_symptom_update_aud
    BEFORE UPDATE ON ${flyway:database}.stm_symptom
    FOR EACH ROW
    set NEW.stm_modification_date=now(),
		NEW.stm_modified_by=USER();

create table ${flyway:database}.cby_comorbidity
(
    cby_id_comorbidity int auto_increment primary key,
    cby_name_comorbidity varchar(50) not null,
    cby_comorbidity_description varchar(100) null,
    cby_id_state_data int null default 1,
    cby_creation_date timestamp null,
    cby_created_by varchar(45) null,
    cby_modification_date timestamp null,
    cby_modified_by varchar(45) null,
    foreign key (cby_id_state_data) references ste_state_data(ste_id_state_data)
);

create trigger ${flyway:database}.cby_comorbidity_insert_aud
    BEFORE INSERT ON ${flyway:database}.cby_comorbidity
    FOR EACH ROW
    set NEW.cby_created_by=USER(),
		NEW.cby_creation_date=now();

create trigger ${flyway:database}.cby_comorbidity_update_aud
    BEFORE UPDATE ON ${flyway:database}.cby_comorbidity
    FOR EACH ROW
    set NEW.cby_modification_date=now(),
		NEW.cby_modified_by=USER();

create table ${flyway:database}.cty_country
(
    cty_id_country int auto_increment primary key not null,
    cty_country_name varchar(45) not null,
    cty_creation_date timestamp null,
    cty_created_by varchar(45) null,
    cty_modification_date timestamp null,
    cty_modified_by varchar(45) null
);

create trigger ${flyway:database}.cty_country_insert_aud
    BEFORE INSERT ON ${flyway:database}.cty_country
    FOR EACH ROW
    set NEW.cty_created_by=USER(),
		NEW.cty_creation_date=now();

create trigger ${flyway:database}.cty_country_update_aud
    BEFORE UPDATE ON ${flyway:database}.cty_country
    FOR EACH ROW
    set NEW.cty_modification_date=now(),
		NEW.cty_modified_by=USER();



create table ${flyway:database}.dpt_department
(
    dpt_id_department int auto_increment primary key,
    dpt_name_dapartment varchar(45) not null,
    dpt_id_country int not null,
    dpt_creation_date timestamp null,
    dpt_created_by varchar(45) null,
    dpt_modification_date timestamp null,
    dpt_modified_by varchar(45) null,
    foreign key (dpt_id_country) references cty_country(cty_id_country)
);

create trigger ${flyway:database}.dpt_department_insert_aud
    BEFORE INSERT ON ${flyway:database}.dpt_department
    FOR EACH ROW
    set NEW.dpt_created_by=USER(),
		NEW.dpt_creation_date=now();

create trigger ${flyway:database}.dpt_department_update_aud
    BEFORE UPDATE ON ${flyway:database}.dpt_department
    FOR EACH ROW
    set NEW.dpt_modification_date=now(),
		NEW.dpt_modified_by=USER();





create table ${flyway:database}.spt_state_patient
(
    spt_id_state_patient int auto_increment primary key not null,
    spt_name_state_patient varchar(45) not null,
    spt_id_state_data int null default 1,
    spt_creation_date timestamp null,
    spt_created_by varchar(45) null,
    spt_modification_date timestamp null,
    spt_modified_by varchar(45) null,
    foreign key (spt_id_state_data) references ste_state_data(ste_id_state_data)
);

create trigger ${flyway:database}.spt_state_patient_insert_aud
    BEFORE INSERT ON ${flyway:database}.spt_state_patient
    FOR EACH ROW
    set NEW.spt_created_by=USER(),
		NEW.spt_creation_date=now();

create trigger ${flyway:database}.spt_state_patient_update_aud
    BEFORE UPDATE ON ${flyway:database}.spt_state_patient
    FOR EACH ROW
    set NEW.spt_modification_date=now(),
		NEW.spt_modified_by=USER();



create table ${flyway:database}.pat_patient
(
    pat_id_patient int auto_increment primary key not null,
    pat_first_name varchar(50) not null,
    pat_second_name varchar(50) null,
    pat_first_last_name varchar(50) not null,
    pat_second_last_name varchar(100) null,
    pat_date_birth date not null,
    pat_id_document_type int not null,
    pat_document_number varchar(15) null,
    pat_cellphone_number varchar(50)  null,
    pat_phone_number varchar(50) null,
    pat_reponsible_family varchar(50) null,
    pat_responsible_family_phone_number varchar(50) null,
    pat_id_department int not null,
    pat_id_country int not null,
    pat_id_patient_sex int not null,
    pat_creation_date timestamp null,
    pat_created_by varchar(45) null,
    pat_modification_date timestamp null,
    pat_modified_by varchar(45) null,
    foreign key (pat_id_department) references dpt_department(dpt_id_department),
    foreign key (pat_id_country) references cty_country(cty_id_country),
    foreign key (pat_id_patient_sex) references sex_patient(spt_id_sex),
    foreign key (pat_id_document_type) references dct_document_type(dct_id_document_type)
);

create trigger ${flyway:database}.pat_patient_insert_aud
    BEFORE INSERT ON ${flyway:database}.pat_patient
    FOR EACH ROW
    set NEW.pat_created_by=USER(),
		NEW.pat_creation_date=now();

create trigger ${flyway:database}.pat_patient_update_aud
    BEFORE UPDATE ON ${flyway:database}.pat_patient
    FOR EACH ROW
    set NEW.pat_modification_date=now(),
		NEW.pat_modified_by=USER();

create table ${flyway:database}.pfl_patient_file
(
    pfl_id_patient_file int auto_increment primary key not null,
    pfl_admission_date datetime not null,
    pfl_high tinyint null default false,
    pfl_high_date date default null,
    pfl_pregnant tinyint null default false,
    pfl_low tinyint null default false,
    pfl_low_date date default null,
    pfl_id_patient int not null,
    pfl_id_state_patient int not null default 1,
    pft_status_bot boolean default false,
    pfl_creation_date timestamp null,
    pfl_created_by varchar(45) null,
    pfl_modification_date timestamp null,
    pfl_modified_by varchar(45) null,
    foreign key (pfl_id_patient) references pat_patient(pat_id_patient),
    foreign key (pfl_id_state_patient) references spt_state_patient(spt_id_state_patient)
);

create trigger ${flyway:database}.pfl_patient_file_insert_aud
    BEFORE INSERT ON ${flyway:database}.pfl_patient_file
    FOR EACH ROW
    set NEW.pfl_created_by=USER(),
		NEW.pfl_creation_date=now();

create trigger ${flyway:database}.pfl_patient_file_update_aud
    BEFORE UPDATE ON ${flyway:database}.pfl_patient_file
    FOR EACH ROW
    set NEW.pfl_modification_date=now(),
		NEW.pfl_modified_by=USER();


create table ${flyway:database}.fhs_file_has_sympton
(
    fhs_id_file_has_sympton int auto_increment primary key,
    fhs_id_symptom int not null,
    fhs_id_patient_file int not null,
    fhs_creation_date timestamp null,
    fhs_created_by varchar(45) null,
    fhs_modification_date timestamp null,
    fhs_modified_by varchar(45) null,
    foreign key (fhs_id_symptom) references stm_symptom(stm_id_sympton),
    foreign key (fhs_id_patient_file) references pfl_patient_file(pfl_id_patient_file)

);

create trigger ${flyway:database}.fhs_file_has_sympton_insert_aud
    BEFORE INSERT ON ${flyway:database}.fhs_file_has_sympton
    FOR EACH ROW
    set NEW.fhs_created_by=USER(),
		NEW.fhs_creation_date=now();

create trigger ${flyway:database}.fhs_file_has_sympton_update_aud
    BEFORE UPDATE ON ${flyway:database}.fhs_file_has_sympton
    FOR EACH ROW
    set NEW.fhs_modification_date=now(),
		NEW.fhs_modified_by=USER();

create table ${flyway:database}.fhc_file_has_cormobility
(
    fhc_id_file_has_cormobility int auto_increment primary key,
    fhc_id_patient_file int not null,
    fhc_id_conmorbilities int not null,
    fhc_creation_date timestamp null,
    fhc_created_by varchar(45) null,
    fhc_modification_date timestamp null,
    fhc_modified_by varchar(45) null,
    foreign key (fhc_id_patient_file) references pfl_patient_file(pfl_id_patient_file),
    foreign key (fhc_id_conmorbilities) references cby_comorbidity(cby_id_comorbidity)
);

create trigger ${flyway:database}.fhc_file_has_cormobility_insert_aud
    BEFORE INSERT ON ${flyway:database}.fhc_file_has_cormobility
    FOR EACH ROW
    set NEW.fhc_created_by=USER(),
		NEW.fhc_creation_date=now();

create trigger ${flyway:database}.fhc_file_has_cormobility_update_aud
    BEFORE UPDATE ON ${flyway:database}.fhc_file_has_cormobility
    FOR EACH ROW
    set NEW.fhc_modification_date=now(),
		NEW.fhc_modified_by=USER();