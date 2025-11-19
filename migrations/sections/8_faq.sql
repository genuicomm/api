INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'FAQ',
    8,
    'faq',
    '[
      {
        "key": "title",
        "label": "Title",
        "type": "text",
        "value": "Pertanyaan Umum",
        "is_visible": true
      },
      {
        "key": "subtitle",
        "label": "Sub Title",
        "type": "text",
        "value": "FAQ",
        "is_visible": true
      },
      {
        "key": "description",
        "label": "Description",
        "type": "textarea",
        "value": "Temukan jawaban atas pertanyaan yang sering diajukan",
        "is_visible": true
      },
      {
        "key": "faq_items",
        "label": "FAQ List",
        "type": "array",
        "value": [
          {
            "key": "faq1",
            "question": "Siapa yang cocok mengikuti pelatihan ini?",
            "answer": "Program kami cocok untuk profesional, entrepreneur, mahasiswa, atau siapa saja yang ingin meningkatkan kemampuan public speaking. Kami memiliki program yang disesuaikan dengan berbagai level kemampuan."
          },
          {
            "key": "faq2",
            "question": "Berapa lama durasi pelatihan?",
            "answer": "Durasi pelatihan bervariasi tergantung program yang dipilih, mulai dari 8 hingga 16 sesi. Setiap sesi berlangsung selama 2 jam dan dapat disesuaikan dengan jadwal Anda."
          },
          {
            "key": "faq3",
            "question": "Apakah ada sesi konsultasi pribadi?",
            "answer": "Ya, setiap program mencakup sesi konsultasi pribadi dengan mentor. Program Professional dan Master bahkan menyediakan lebih banyak sesi one-on-one coaching."
          }
        ],
        "is_visible": true
      }
    ]'::jsonb
);