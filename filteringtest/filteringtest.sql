CREATE TABLE t_policy (
	policy_number integer,
	policy_submit_date date,
	premium money,
	discount integer,
	commission money,
	client_number varchar(20) unique,
	agent_code varchar(20) unique,
	policy_status varchar(25),
	policy_due_date date,
	constraint t_policy_pk primary key (policy_number, client_number, agent_code)
)

CREATE TABLE t_agent(
	agent_code varchar(20),
	agent_name varchar(50),
	agent_office varchar(30),
	basic_commission money,
	foreign key (agent_code) references t_policy(agent_code)
)

CREATE TABLE t_client(
	client_number varchar(20),
	client_name varchar(50),
	birth_date date,
	foreign key (client_number) references t_policy(client_number)
)

INSERT INTO t_policy (policy_number, policy_submit_date, premium, discount, commission, client_number, agent_code, policy_status, policy_due_date) VALUES 
(001, '2018-1-5', 18600000, 10, 930000, 'CL001', 'AG001', 'INFORCE', null),
(002, '2018-1-5', 2500000, 10, 125000, 'CL002', 'AG002', 'INFORCE', null),
(003, '2018-1-10', 2500000, 10, 125000, 'CL003', 'AG003', 'TERMINATE', null),
(004, '2018-1-25', 4000000, 10, 200000, 'CL003', 'AG002', 'PROPOSAL', null),
(005, '2018-1-25', 2625000, 10, 131250, 'CL004', 'AG002', 'PROPOSAL', null)

INSERT INTO t_agent (agent_code, agent_name, agent_office, basic_commission) VALUES
('AG001', 'LIDYA', 'JAKARTA', NULL),
('AG002', 'RUDI', 'BALI', NULL),
('AG003', 'DICKY', 'BALI', NULL),
('AG004', 'DAVID', 'SURABAYA', NULL),
('AG005', 'FARIZ', 'SURABAYA', NULL)

INSERT INTO t_client (client_number, client_name, birth_date) VALUES
('CL001', 'WAYNE ROONEY', '1956-5-11'),
('CL002', 'RYAN GIGGS', '1972-9-3'),
('CL003', 'DAVID BECKHAM', '1985-9-3'),
('CL004', 'MICHAEL OWEN', '2012-9-3')

SELECT * FROM t_policy
SELECT * FROM t_agent
SELECT * FROM t_client

-- Soal nomor 1 SQL 
-- A
-- Tampilkan data polis yang disubmit setelah tanggal 15 januari 2018 berdasarkan client yang lahir di bulan SEPTEMBER
SELECT * FROM t_policy a
	LEFT JOIN t_client b
	ON a.client_number = b.client_number
	WHERE policy_submit_date > '2018-01-15' AND EXTRACT(MONTH FROM birth_date) = 9;
	
-- B
-- Tampilkan data polis dengan status INFORCE, yang agentnya berkantor di "Jakarta"
SELECT * FROM t_policy a
	LEFT JOIN t_agent b
	ON a.agent_code = b.agent_code
	WHERE policy_status = 'INFORCE' AND agent_office = 'JAKARTA';


-- C
-- Hitunglah kolom basic commission pada table T_Agent dengan rumus:
-- nilai (COMMISSION/PREMIUM * 100 dari table T_Policy
UPDATE t_agent a
	SET basic_commission = (b.commission/b.premium) * 100 :: MONEY
	FROM t_policy b
	WHERE basic_commission IS NULL

-- D
-- Isilah kolom Policy_Due_Date dari table T_Policy, dengan kondisi tanggal akhir bulan dari 30 hari setelah tanggal Policy_Submit_Date
-- DATE TRUNC digunakan untuk memotong tanggal menjadi awal bulan dari 30 hari setelah tanggal "Policy_Submit_Date".
UPDATE t_policy
    SET policy_due_date = DATE_TRUNC('MONTH', policy_submit_date + INTERVAL '30 days') + INTERVAL '1 MONTH' - INTERVAL '1 DAY';

-- E
-- Tampilkan semua data produksi agent yang nilai premiumnya setelah dipotong DISCOUNT dibawah SATU JUTA (diurutkan dari nilai terkecil)
SELECT a.agent_code, a.agent_name, a.agent_office, p.premium - cast(p.premium * 0.1 AS MONEY) AS produksi_agent
    FROM t_agent AS a 
    LEFT JOIN t_policy as p ON p.agent_code = a.agent_code
    WHERE cast(p.premium * 0.1 AS MONEY) < cast(1000000 AS MONEY)
    ORDER BY produksi_agent ASC