INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Service',
    4,
    'service',
    '[
  {
    "key": "title",
    "type": "text",
    "label": "Title",
    "value": "Layanan Kami",
    "is_visible": true  
  },
  {
    "key": "subtitle",
    "type": "text",
    "label": "Sub Title",
    "value": "OUR SERVICES",
    "is_visible": true
  },
  {
    "key": "description",
    "type": "textarea",
    "label": "Description",
    "value": "Program pelatihan kami dirancang dengan berbagai tingkatan kelas yang disesuaikan dengan tingkat kesulitan dan kebutuhan peserta.",
    "is_visible": true
  },
  {
    "key": "services",
    "type": "array",
    "label": "Services",
    "value": [
      {
        "key": "public_speaking",
        "image": "https://public.readdy.ai/ai/img_res/cb2ccd39b5fcad5d842ca5ea7ff85e86.jpg",
        "icon": "ri-presentation-line",
        "title": "Public Speaking Training",
        "description": "Program pelatihan komprehensif untuk meningkatkan kemampuan berbicara di depan umum dengan metode yang terstruktur dan personal.",
        "features": [
          "Profil Komunikasi Personal",
          "Optimalisasi Napas, Vokal & Postur",
          "Strategi Penyampaian Konten",
          "Optimalisasi Gestur"
        ],
        "detail": "Kelas terbatas 15-20 peserta dengan kombinasi paparan teori, latihan praktis, dan simulasi. Tersedia untuk individu dewasa (min. 18 tahun) dari berbagai latar belakang.",
        "button": {
          "label": "Pelajari Lebih Lanjut",
          "url": "#"
        }
      },
      {
        "key": "corporate_training",
        "image": "https://public.readdy.ai/ai/img_res/fa295f9f3ad7b0e171d89cdb74bd1773.jpg",
        "icon": "ri-building-line",
        "title": "Corporate Training",
        "description": "Program khusus untuk meningkatkan kemampuan komunikasi tim perusahaan yang disesuaikan dengan kebutuhan organisasi Anda.",
        "features": [
          "Presentasi Bisnis Efektif",
          "Komunikasi Tim & Leadership",
          "Simulasi Pola Komunikasi Publik",
          "Strategi Komunikasi Korporat"
        ],
        "detail": "Workshop tatap muka di lokasi perusahaan Anda dengan kurikulum yang disesuaikan. Pelatihan meliputi aktivitas kelompok, latihan praktis dan evaluasi intensif.",
        "button": {
          "label": "Pelajari Lebih Lanjut",
          "url": "#"
        }
      },
      {
        "key": "personal_coaching",
        "image": "https://public.readdy.ai/ai/img_res/abc8f100643d91b09b96a6d897844849.jpg",
        "icon": "ri-user-voice-line",
        "title": "Personal Coaching",
        "description": "Bimbingan one-on-one untuk pengembangan kemampuan komunikasi yang lebih personal dan disesuaikan dengan kebutuhan spesifik Anda.",
        "features": [
          "Analisis Profil Komunikasi Mendalam",
          "Latihan Vokal & Gestur Intensif",
          "Pengembangan Konten Personal",
          "Evaluasi & Feedback Berkala"
        ],
        "detail": "Sesi intensif dengan instruktur berpengalaman kami, termasuk Karwata Widjaja. Program dibuat khusus untuk mengatasi tantangan komunikasi spesifik yang Anda hadapi.",
        "button": {
          "label": "Pelajari Lebih Lanjut",
          "url": "#"
        }
      }
    ],
    "is_visible": true
  },
  {
    "key": "methods",
    "type": "array",
    "label": "Metode Pembelajaran",
    "value": [
      {
        "icon": "ri-discuss-line",
        "title": "Paparan dan Diskusi",
        "description": "Penyampaian konsep dasar dan diskusi interaktif untuk pemahaman mendalam"
      },
      {
        "icon": "ri-voice-recognition-line",
        "title": "Latihan Praktis",
        "description": "Latihan Vokal, Gestur dan Storytelling yang intensif untuk pengembangan teknis"
      },
      {
        "icon": "ri-heart-pulse-line",
        "title": "Senam Pernafasan",
        "description": "Teknik pernapasan dan postur yang mendukung performa berbicara di depan publik"
      },
      {
        "icon": "ri-team-line",
        "title": "Simulasi & Praktek",
        "description": "Simulasi serta praktek individu dan kelompok dalam konteks presentasi nyata"
      }
    ],
    "is_visible": true
  }
]'::jsonb
);