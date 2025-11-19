INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'About',
    2,
    'about',
    '[
      {
        "key": "title",
        "label": "Title",
        "type": "text",
        "value": "Tentang Kami",
        "is_visible": true
      },
      {
        "key": "subtitle",
        "label": "Sub Title",
        "type": "text",
        "value": "About Us",
        "is_visible": true
      },
      {
        "key": "description",
        "label": "Description",
        "type": "textarea",
        "value": "Genuicomm adalah perusahaan yang berfokus pada pengembangan keterampilan komunikasi dan public speaking melalui program pelatihan yang telah diselenggarakan sejak 2019.",
        "is_visible": true
      },
      {
        "key": "company_info",
        "label": "Company Info",
        "type": "textarea",
        "value": "Kami menghadirkan pendekatan yang efektif dan inovatif, dimulai dengan mengenali profil komunikasi diri, kemudian mengoptimalkan latihan vokal, gestur, serta materi komunikasi lainnya. Program pelatihan kami dirancang dengan berbagai tingkatan kelas yang disesuaikan dengan tingkat kesulitan dan kebutuhan peserta.",
        "is_visible": true
      },
      {
        "key": "stats",
        "label": "Stats",
        "type": "array",
        "value": [
          {
            "key": "certified_participants",
            "label": "Peserta Tersertifikasi",
            "value": "5000+",
            "icon": "ri-award-line",
            "label_editable": true
          },
          {
            "key": "company_partners",
            "label": "Perusahaan Partner",
            "value": "200+",
            "icon": "ri-building-line",
            "label_editable": true
          },
          {
            "key": "experienced_trainers",
            "label": "Trainer Berpengalaman",
            "value": "50+",
            "icon": "ri-team-line",
            "label_editable": true
          },
          {
            "key": "years_experience",
            "label": "Tahun Pengalaman",
            "value": "5+",
            "icon": "ri-global-line",
            "label_editable": true
          }
        ],
        "is_visible": true
      },
      {
        "key": "image_1",
        "label": "Gambar 1",
        "type": "file",
        "value": "https://public.readdy.ai/ai/img_res/717bb5c8ab0dbdc8f81f4c71eb2b26e8.jpg",
        "is_visible": true
      },
      {
        "key": "values_title",
        "label": "Nilai-Nilai Kami",
        "type": "textarea",
        "value": "Kami percaya setiap individu memiliki potensi besar yang dapat dikembangkan. Dengan pendekatan yang tepat, setiap orang dapat meningkatkan keterampilan komunikasi dan public speaking mereka untuk mencapai versi terbaik dari dirinya.",
        "is_visible": true
      },
      {
        "key": "values",
        "label": "Nilai-Nilai Utama",
        "type": "array",
        "value": [
          {
            "key": "confident_and_authentic",
            "label": "Percaya Diri & Autentik - Mendorong komunikasi yang jujur dan sesuai karakter",
            "icon": "ri-user-heart-line",
            "label_editable": true
          },
          {
            "key": "continuous_learning",
            "label": "Pembelajaran Berkelanjutan - Metode berbasis riset dan tren terbaru",
            "icon": "ri-book-open-line",
            "label_editable": true
          },
          {
            "key": "professionalism",
            "label": "Profesionalisme & Integritas - Standar tinggi dan etika dalam pelatihan",
            "icon": "ri-shield-check-line",
            "label_editable": true
          },
          {
            "key": "positive_impact",
            "label": "Perubahan & Dampak Positif - Menciptakan transformasi nyata",
            "icon": "ri-seedling-line",
            "label_editable": true
          }
        ],
        "is_visible": true
      },
      {
        "key": "image_2",
        "label": "Gambar 2",
        "type": "file",
        "value": "https://public.readdy.ai/ai/img_res/47191d11c61ea1ef6174c6bd43f3c249.jpg",
        "is_visible": true
      }
    ]'::jsonb
);