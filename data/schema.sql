-- Doctors table
CREATE TABLE IF NOT EXISTS doctors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    name TEXT NOT NULL,
    skill TEXT NOT NULL,
    title TEXT NOT NULL,
    location TEXT NOT NULL
);

-- Patients table
CREATE TABLE IF NOT EXISTS patients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    doctor_id INTEGER,
    name TEXT NOT NULL,
    age INTEGER NOT NULL,
    contact TEXT NOT NULL,
    district TEXT NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctors(id)
);

-- Appointments table
CREATE TABLE IF NOT EXISTS appointments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    doctor_name TEXT NOT NULL,
    patient_name TEXT NOT NULL,
    time DATETIME DEFAULT CURRENT_TIMESTAMP,
    type TEXT NOT NULL,
    platform TEXT NOT NULL,
    FOREIGN KEY (doctor_name) REFERENCES doctors(name),
    FOREIGN KEY (patient_name) REFERENCES patients(name)
);

-- Diagnoses table
CREATE TABLE IF NOT EXISTS diagnoses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    symptoms TEXT NOT NULL,
    treatments TEXT NOT NULL,
    specialty TEXT NOT NULL,
    FOREIGN KEY (specialty) REFERENCES doctors(skill)
);
