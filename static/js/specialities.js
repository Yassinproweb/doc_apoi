document.addEventListener("DOMContentLoaded", () => {
  const specialities = {
    "allergy": {
      "symptoms": "Sneezing, Itchy eyes, Runny nose, Cough, Wheezing, Rash, Swelling, Hives, Shortness of breath, Nasal congestion, Postnasal drip, Itchy throat, Watery eyes, Fatigue, Headache.",
      "procedures": "Skin prick testing, Patch testing, Serum IgE testing, Allergen avoidance counseling, Immunotherapy (allergy shots), Epinephrine auto-injector training, Nasal corticosteroids, Antihistamines, Bronchodilator therapy, Environmental control."
    },
    "anaesthesiology": {
      "symptoms": "Preoperative anxiety, Pain, Nausea, Vomiting, Breathing difficulty, Hypotension, Bradycardia, Drowsiness, Chills, Confusion, Muscle rigidity, Tachycardia, Sweating, Tremor, Delayed recovery.",
      "procedures": "General anesthesia, Regional anesthesia (epidural, spinal, nerve block), Sedation, Airway management, Intubation, Pain control, Ventilator management, Preoperative evaluation, Intraoperative monitoring, Post-anesthesia care."
    },
    "audiology": {
      "symptoms": "Hearing loss, Tinnitus, Vertigo, Dizziness, Ear fullness, Ear pain, Balance disturbances, Noise sensitivity, Speech perception difficulty, Sound distortion, Pressure sensation, Ear drainage, Autophony.",
      "procedures": "Pure tone audiometry, Tympanometry, Otoacoustic emissions testing, Auditory brainstem response, Hearing aid fitting, Cochlear implant programming, Vestibular rehabilitation, Acoustic reflex testing, Balance assessment."
    },
    "cardiology": {
      "symptoms": "Chest pain, Palpitations, Dyspnea, Orthopnea, Fatigue, Edema, Dizziness, Syncope, Cyanosis, Tachycardia, Bradycardia, Claudication, Heart murmur, Chest tightness, Exercise intolerance, Weakness, Fainting, Cough, Cold extremities.",
      "procedures": "ECG, Echocardiography, Stress testing, Cardiac catheterization, Angioplasty, Stent placement, Pacemaker insertion, Valve replacement, Coronary bypass surgery, Holter monitoring, Cardioversion, Cardiac rehabilitation, Thrombolysis, Cardiac MRI."
    },
    "dentistry": {
      "symptoms": "Toothache, Bleeding gums, Jaw pain, Bad breath, Tooth decay, Swollen gums, Loose teeth, Tooth sensitivity, Mouth ulcers, Jaw stiffness, Gum recession, Pain on chewing, Swelling, Infection, Cracked teeth.",
      "procedures": "Dental cleaning, Tooth extraction, Root canal treatment, Filling, Crown placement, Scaling and polishing, Orthodontic treatment, Dental X-rays, Teeth whitening, Implants, Gum surgery, Bridge and veneer fitting."
    },
    "dermatology": {
      "symptoms": "Rash, Itching, Redness, Dry skin, Blisters, Acne, Hair loss, Moles, Skin pain, Urticaria, Scaling, Pigmentation changes, Nail changes, Psoriasis, Ulcers, Eczema, Sun sensitivity.",
      "procedures": "Skin biopsy, Cryotherapy, Laser therapy, Phototherapy, Excision of lesions, Chemical peels, Mohs surgery, Topical therapy, Intralesional injections, Scar revision, Mole removal, Skin grafting."
    },
    "gastroenterology": {
      "symptoms": "Abdominal pain, Bloating, Nausea, Vomiting, Heartburn, Diarrhea, Constipation, Jaundice, Weight loss, Loss of appetite, Dysphagia, Hematemesis, Melena, Rectal bleeding, Indigestion, Flatulence, Gas, Abdominal distension.",
      "procedures": "Endoscopy, Colonoscopy, ERCP, Liver biopsy, Capsule endoscopy, PEG tube placement, Paracentesis, Polypectomy, Sigmoidoscopy, Hemostasis, Variceal ligation, GI surgery, Manometry, Biliary drainage."
    },
    "geriatrics": {
      "symptoms": "Memory loss, Confusion, Fatigue, Falls, Incontinence, Depression, Tremor, Frailty, Weakness, Sleep disorders, Malnutrition, Polypharmacy effects, Cognitive decline, Weight loss, Mobility issues.",
      "procedures": "Comprehensive geriatric assessment, Medication review, Cognitive testing, Mobility evaluation, Fall risk assessment, Nutritional counseling, Vision and hearing screening, Blood pressure control, Dementia management."
    },
    "gynaecology": {
      "symptoms": "Pelvic pain, Irregular menstruation, Vaginal discharge, Menorrhagia, Dysmenorrhea, Infertility, Hot flashes, Dyspareunia, Itching, Postmenopausal bleeding, Vaginal dryness, Lower back pain, Pelvic pressure, Spotting.",
      "procedures": "Pap smear, Pelvic ultrasound, Hysteroscopy, Laparoscopy, Colposcopy, Endometrial biopsy, IUD insertion, Oophorectomy, Hysterectomy, C-section, D&C, Myomectomy, Fertility treatments."
    },
    "nephrology": {
      "symptoms": "Edema, Hypertension, Fatigue, Decreased urine output, Foamy urine, Hematuria, Nausea, Vomiting, Confusion, Pruritus, Weakness, Shortness of breath, Weight gain, Leg swelling.",
      "procedures": "Dialysis, Kidney biopsy, Urinalysis, Electrolyte management, Renal ultrasound, Vascular access creation, Fluid management, Renal function tests, Kidney transplant preparation."
    },
    "neurology": {
      "symptoms": "Headache, Dizziness, Weakness, Numbness, Seizures, Tremor, Memory loss, Ataxia, Syncope, Visual disturbance, Slurred speech, Paresthesia, Gait imbalance, Paralysis, Vertigo.",
      "procedures": "EEG, EMG, Lumbar puncture, MRI brain, CT scan, Carotid Doppler, Botox injection, Nerve conduction study, Thrombectomy, Deep brain stimulation, Spinal tap, Neuroimaging analysis."
    },
    "nutrition": {
      "symptoms": "Weight loss, Fatigue, Weakness, Malnutrition, Hair loss, Brittle nails, Loss of appetite, Dry skin, Cramping, Dizziness, Fainting, Low energy, Poor concentration.",
      "procedures": "Dietary assessment, Meal planning, Nutritional counseling, BMI measurement, Micronutrient testing, Supplement therapy, Weight management, Caloric intake tracking, Diet modification."
    },
    "oncology": {
      "symptoms": "Fatigue, Weight loss, Pain, Lump or mass, Bleeding, Cough, Jaundice, Night sweats, Fever, Anemia, Bone pain, Nausea, Vomiting, Weakness, Swelling, Pallor.",
      "procedures": "Chemotherapy, Radiation therapy, Tumor biopsy, Surgical resection, Targeted therapy, Immunotherapy, Bone marrow transplant, Palliative care, PET-CT imaging, Genetic testing."
    },
    "ophthalmology": {
      "symptoms": "Blurry vision, Eye pain, Redness, Photophobia, Tearing, Floaters, Double vision, Loss of vision, Swelling, Discharge, Itching, Burning, Haloes, Dry eyes.",
      "procedures": "Visual acuity testing, Fundoscopy, Tonometry, Cataract surgery, Laser correction, Retinal detachment repair, Intravitreal injection, Glaucoma management, Corneal transplant."
    },
    "orthodontics": {
      "symptoms": "Crooked teeth, Bite misalignment, Jaw pain, Speech difficulty, Difficulty chewing, Tooth crowding, Overbite, Underbite, Teeth grinding, Facial imbalance.",
      "procedures": "Braces placement, Retainers, Teeth alignment, Jaw repositioning, Orthognathic surgery, Invisalign fitting, Dental arch expansion, Functional appliances."
    },
    "orthopaedics": {
      "symptoms": "Joint pain, Fracture, Swelling, Stiffness, Limited mobility, Muscle weakness, Back pain, Joint deformity, Tenderness, Locking, Numbness, Inflammation.",
      "procedures": "Fracture fixation, Joint replacement, Arthroscopy, Casting and splinting, Physical therapy, Spinal surgery, Ligament reconstruction, Bone grafting, Orthotic fitting."
    },
    "otolaryngology": {
      "symptoms": "Hearing loss, Tinnitus, Otalgia, Vertigo, Nasal congestion, Rhinorrhea, Epistaxis, PND, Facial pain, Anosmia, Sore throat, Dysphagia, Hoarseness, Snoring, Apnea, Neck mass.",
      "procedures": "Cerumen removal, PE tubes, INCS for AR/CRS, FESS, Septoplasty, Abx for tonsillitis, Tonsillectomy, Voice therapy, CPAP for OSA, Thyroidectomy, LN biopsy."
    },
    "pathology": {
      "symptoms": "Abnormal lab tests, Cytologic anomalies, Blood disorders, Tumor detection, Cell abnormalities, Inflammatory patterns, Infection evidence, Tissue necrosis.",
      "procedures": "Histopathology, Cytopathology, Molecular diagnostics, Immunohistochemistry, Autopsy, Frozen section, Hematopathology, Microbiologic culture, Genetic analysis."
    },
    "pediatrics": {
      "symptoms": "Fever, Cough, Rash, Vomiting, Diarrhea, Growth delay, Seizures, Ear pain, Lethargy, Crying, Abdominal pain, Failure to thrive, Poor feeding, Dehydration.",
      "procedures": "Vaccination, Growth monitoring, Developmental screening, Lumbar puncture, IV access, Newborn resuscitation, Pediatric endoscopy, Imaging, Nutritional support."
    },
    "pharmacy": {
      "symptoms": "Medication side effects, Drug interactions, Adverse reactions, Overdose, Polypharmacy, Allergy symptoms, Non-compliance, Withdrawal signs.",
      "procedures": "Medication therapy management, Pharmacovigilance, Compounding, Prescription review, Patient counseling, Dosing adjustment, Medication reconciliation."
    },
    "psychiatry": {
      "symptoms": "Depression, Anxiety, Insomnia, Hallucinations, Delusions, Mood swings, Suicidal thoughts, Irritability, Lack of motivation, Apathy, Memory loss, Obsessions.",
      "procedures": "Psychotherapy, Medication management, Cognitive behavioral therapy, Electroconvulsive therapy, Group therapy, Crisis intervention, Behavioral assessment."
    },
    "pulmonology": {
      "symptoms": "Cough, Dyspnea, Chest pain, Hemoptysis, Wheezing, Fatigue, Sleep apnea, Cyanosis, Sputum production, Fever, Breathlessness, Chronic cough.",
      "procedures": "Spirometry, Bronchoscopy, Thoracentesis, Mechanical ventilation, Lung biopsy, Chest physiotherapy, Pleural drainage, Pulmonary rehab, Sleep studies."
    },
    "radiology": {
      "symptoms": "Imaging abnormalities, Pain, Swelling, Trauma assessment, Lesion detection, Mass identification, Organ enlargement, Bone deformity, Internal bleeding.",
      "procedures": "X-ray, CT scan, MRI, Ultrasound, Fluoroscopy, Interventional radiology, PET scan, Angiography, Image-guided biopsy, Contrast imaging."
    },
    "rheumatology": {
      "symptoms": "Joint pain, Swelling, Morning stiffness, Fatigue, Fever, Rash, Muscle aches, Deformities, Weakness, Limited mobility, Tenderness, Pain flare-ups.",
      "procedures": "Joint aspiration, Synovial biopsy, Autoantibody testing, Steroid injections, Disease-modifying therapy, Pain management, Physical therapy, Immunosuppressive treatment."
    },
    "urology": {
      "symptoms": "Dysuria, Hematuria, Incontinence, Flank pain, Pelvic pain, Erectile dysfunction, Nocturia, Urinary retention, Urgency, Frequency, Cloudy urine.",
      "procedures": "Cystoscopy, Prostate exam, Lithotripsy, Catheterization, TURP, Vasectomy, Urodynamic testing, Ureteroscopy, Prostate biopsy, Bladder surgery."
    }
  },
    docSkill = document.getElementById("skill-doc"),
    symptoms = document.getElementById("symptoms"),
    procedures = document.getElementById("procedures");

  if (!docSkill || !symptoms || !procedures) return;

  const skill = docSkill.textContent.trim().toLowerCase();
  if (specialities[skill]) {
    symptoms.textContent = specialities[skill].symptoms;
    procedures.textContent = specialities[skill].procedures;
  } else {
    symptoms.textContent = "Symptoms data not available!";
    procedures.textContent = "Procedures data not available!"
  }
});
