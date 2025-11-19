INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'CTA',
    9,
    'cta',
    '[
      {
        "key": "title",
        "label": "Judul",
        "type": "text",
        "value": "Siap Untuk Memaksimalkan Potensi Anda?",
        "is_visible": true
      },
      {
        "key": "description",
        "label": "Deskripsi",
        "type": "textarea",
        "value": "Bergabunglah dengan program komunikasi dan public speaking kami yang telah membantu ribuan individu dan organisasi sejak 2019. Kelas terbatas hanya 15-20 peserta untuk memastikan perhatian intensif.",
        "is_visible": true
      },
      {
        "key": "buttons",
        "label": "Hubungi Kami",
        "type": "button",
        "value": "/contact",
        "is_visible": true
      }
    ]'::jsonb
);