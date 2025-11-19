INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Testimoni',
    6,
    'testimoni',
    '[
      {
        "key": "title",
        "label": "Title",
        "value": "Apa kata mereka tentang kami",
        "type": "text",
        "is_visible": true
      },
      {
        "key": "subtitle",
        "label": "Sub Title",
        "value": "Testimonials",
        "type": "text",
        "is_visible": true
      },
      {
        "key": "description",
        "label": "Description",
        "value": "Testimoni nyata dari mereka yang telah merasakan manfaat program kami.",
        "type": "textarea",
        "is_visible": true
      },
      {
        "key": "video_testimonial",
        "label": "Video Testimonial",
        "video_url": "https://www.instagram.com/reel/DJEZ7LfyChJ/",
        "name": "John Doe",
        "role": "CEO Company",
        "quote": "Sangat puas dengan layanan yang diberikan. Tim sangat profesional dan hasilnya sesuai ekspektasi.",
        "type": "reels",
        "is_visible": true
      },
      {
        "key": "text_testimonials",
        "label": "Testimoni Teks",
        "value": [
          {
            "key": "t1",
            "photo": "https://public.readdy.ai/ai/img_res/29f20197799a1db68144e23446e9a2ad.jpg",
            "name": "Dewi Pratiwi",
            "role": "Marketing Manager",
            "quote": "Program ini benar-benar mengubah cara saya berkomunikasi. Sekarang saya lebih percaya diri dalam presentasi dan meeting dengan klien."
          },
          {
            "key": "t2",
            "photo": "https://public.readdy.ai/ai/img_res/aca69b75beded4e419f1fc4c624ebbb8.jpg",
            "name": "Budi Santoso",
            "role": "CEO Startup",
            "quote": "Mentor yang sangat berpengalaman dan metode yang efektif. Saya bisa langsung praktekkan ilmunya di pitch deck startup saya."
          },
          {
            "key": "t3",
            "photo": "https://public.readdy.ai/ai/img_res/d951b010e97a792b95cba233bcb10aff.jpg",
            "name": "Rina Wijaya",
            "role": "Public Relations",
            "quote": "Materi yang komprehensif dan praktek yang intensif membuat saya lebih siap menghadapi media dan public speaking."
          }
        ],
        "type": "array",
        "is_visible": true
      }
    ]'::jsonb
);