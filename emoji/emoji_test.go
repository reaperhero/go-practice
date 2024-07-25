package emoji

import (
	"encoding/json"
	"fmt"
	"github.com/kyokomi/emoji/v2"
	"io"
	"os"
	"testing"
)

//😀 :grinning: 😁 :grin: 😂 :joy: 🤣 :rofl: 😃 :smiley: 😄 :grin: 😅 :sweat_smile: 😆 :laughing: 😉 :wink:
//😊 :blush: 😋 :yum: 😎 :sunglasses: 😍 :heart_eyes: 😘 :kissing_heart: 😗 :kissing: 😙 :kissing_smiling_eyes: 😚 :kissing_closed_eyes: ☺ :relaxed:
//🙂 :slight_smile: 🤗 :hugging: 🤩 :star_struck: 🤔 :thinking: 🤨 :face_with_raised_eyebrow: 😐 :neutral_face: 😑 :expressionless: 😶 :no_mouth: 🙄 :rolling_eyes:
//😏 :smirk: 😣 :persevere: 😥 :disappointed_relieved: 😮 :open_mouth: 🤐 :zipper_mouth: 😯 :hushed: 😪 :sleepy: 😫 :tired_face: 😴 :sleeping:
//😌 :relieved: 😛 :stuck_out_tongue: 😜 :stuck_out_tongue_winking_eye: 😝 :stuck_out_tongue_closed_eyes: 🤤 :drooling_face: 😒 :unamused: 😓 :sweat: 😔 :pensive: 😕 :confused:
//🙃 :upside_down: 🤑 :money_mouth: 😲 :astonished: ☹ :frowning2: 🙁 :slight_frown: 😖 :confounded: 😞 :disappointed: 😟 :worried: 😤 :triumph:
//😢 :cry: 😭 :sob: 😦 :frowning: 😧 :anguished: 😨 :fearful: 😩 :weary: 🤯 :exploding_head: 😬 :grimacing: 😰 :cold_sweat:
//😱 :scream: 😳 :flushed: 🤪 :crazy_face: 😵 :dizzy_face: 😡 :rage: 😠 :angry: 🤬 :face_with_symbols_over_mouth: 😷 :mask: 🤒 :thermometer_face:
//🤕 :head_bandage: 🤢 :nauseated_face: 🤮 :face_vomiting: 🤧 :sneezing_face: 😇 :innocent: 🤠 :cowboy: 🤡 :clown: 🤥 :lying_face: 🤫 :shushing_face:
//🤭 :face_with_hand_over_mouth: 🧐 :face_with_monocle: 🤓 :nerd: 😈 :smiling_imp: 👿 :imp: 👹 :japanese_ogre: 👺 :japanese_goblin: 💀 :skull: ☠ :skull_crossbones:
//👻 :ghost: 👽 :alien: 👾 :space_invader: 🤖 :robot: 💩 :poop: 😺 :smiley_cat: 😸 :smile_cat: 😹 :joy_cat: 😻 :heart_eyes_cat:
//😼 :smirk_cat: 😽 :kissing_cat: 🙀 :scream_cat: 😿 :crying_cat_face: 😾 :pouting_cat: 🙈 :see_no_evil: 🙉 :hear_no_evil: 🙊 :speak_no_evil: 👶 :baby:
//👶🏻 :baby_tone1: 👶🏼 :baby_tone2: 👶🏽 :baby_tone3: 👶🏾 :baby_tone4: 👶🏿 :baby_tone5: 🧒 :child: 🧒🏻 :child_tone1: 🧒🏼 :child_tone2: 🧒🏽 :child_tone3:
//🧒🏾 :child_tone4: 🧒🏿 :child_tone5: 👦 :boy: 👦🏻 :boy_tone1: 👦🏼 :boy_tone2: 👦🏽 :boy_tone3: 👦🏾 :boy_tone4: 👦🏿 :boy_tone5: 👧 :girl:
//👧🏻 :girl_tone1: 👧🏼 :girl_tone2: 👧🏽 :girl_tone3: 👧🏾 :girl_tone4: 👧🏿 :girl_tone5: 🧑 :adult: 🧑🏻 :adult_tone1: 🧑🏼 :adult_tone2: 🧑🏽 :adult_tone3:
//🧑🏾 :adult_tone4: 🧑🏿 :adult_tone5: 👨 :man: 👨🏻 :man_tone1: 👨🏼 :man_tone2: 👨🏽 :man_tone3: 👨🏾 :man_tone4: 👨🏿 :man_tone5: 👩 :woman:
//👩🏻 :woman_tone1: 👩🏼 :woman_tone2: 👩🏽 :woman_tone3: 👩🏾 :woman_tone4: 👩🏿 :woman_tone5: 🧓 :older_adult: 🧓🏻 :older_adult_tone1: 🧓🏼 :older_adult_tone2: 🧓🏽 :older_adult_tone3:
//🧓🏾 :older_adult_tone4: 🧓🏿 :older_adult_tone5: 👴 :older_man: 👴🏻 :older_man_tone1: 👴🏼 :older_man_tone2: 👴🏽 :older_man_tone3: 👴🏾 :older_man_tone4: 👴🏿 :older_man_tone5: 👵 :older_woman:
//👵🏻 :older_woman_tone1: 👵🏼 :older_woman_tone2: 👵🏽 :older_woman_tone3: 👵🏾 :older_woman_tone4: 👵🏿 :older_woman_tone5: 👨‍⚕️ :man_health_worker: 👨🏻‍⚕️ :man_health_worker_tone1: 👨🏼‍⚕️ :man_health_worker_tone2: 👨🏽‍⚕️ :man_health_worker_tone3:
//👨🏾‍⚕️ :man_health_worker_tone4: 👨🏿‍⚕️ :man_health_worker_tone5: 👩‍⚕️ :woman_health_worker: 👩🏻‍⚕️ :woman_health_worker_tone1: 👩🏼‍⚕️ :woman_health_worker_tone2: 👩🏽‍⚕️ :woman_health_worker_tone3: 👩🏾‍⚕️ :woman_health_worker_tone4: 👩🏿‍⚕️ :woman_health_worker_tone5: 👨‍🎓 :man_student:
//👨🏻‍🎓 :man_student_tone1: 👨🏼‍🎓 :man_student_tone2: 👨🏽‍🎓 :man_student_tone3: 👨🏾‍🎓 :man_student_tone4: 👨🏿‍🎓 :man_student_tone5: 👩‍🎓 :woman_student: 👩🏻‍🎓 :woman_student_tone1: 👩🏼‍🎓 :woman_student_tone2: 👩🏽‍🎓 :woman_student_tone3:
//👩🏾‍🎓 :woman_student_tone4: 👩🏿‍🎓 :woman_student_tone5: 👨‍🏫 :man_teacher: 👨🏻‍🏫 :man_teacher_tone1: 👨🏼‍🏫 :man_teacher_tone2: 👨🏽‍🏫 :man_teacher_tone3: 👨🏾‍🏫 :man_teacher_tone4: 👨🏿‍🏫 :man_teacher_tone5: 👩‍🏫 :woman_teacher:
//👩🏻‍🏫 :woman_teacher_tone1: 👩🏼‍🏫 :woman_teacher_tone2: 👩🏽‍🏫 :woman_teacher_tone3: 👩🏾‍🏫 :woman_teacher_tone4: 👩🏿‍🏫 :woman_teacher_tone5: 👨‍⚖️ :man_judge: 👨🏻‍⚖️ :man_judge_tone1: 👨🏼‍⚖️ :man_judge_tone2: 👨🏽‍⚖️ :man_judge_tone3:
//👨🏾‍⚖️ :man_judge_tone4: 👨🏿‍⚖️ :man_judge_tone5: 👩‍⚖️ :woman_judge: 👩🏻‍⚖️ :woman_judge_tone1: 👩🏼‍⚖️ :woman_judge_tone2: 👩🏽‍⚖️ :woman_judge_tone3: 👩🏾‍⚖️ :woman_judge_tone4: 👩🏿‍⚖️ :woman_judge_tone5: 👨‍🌾 :man_farmer:
//👨🏻‍🌾 :man_farmer_tone1: 👨🏼‍🌾 :man_farmer_tone2: 👨🏽‍🌾 :man_farmer_tone3: 👨🏾‍🌾 :man_farmer_tone4: 👨🏿‍🌾 :man_farmer_tone5: 👩‍🌾 :woman_farmer: 👩🏻‍🌾 :woman_farmer_tone1: 👩🏼‍🌾 :woman_farmer_tone2: 👩🏽‍🌾 :woman_farmer_tone3:
//👩🏾‍🌾 :woman_farmer_tone4: 👩🏿‍🌾 :woman_farmer_tone5: 👨‍🍳 :man_cook: 👨🏻‍🍳 :man_cook_tone1: 👨🏼‍🍳 :man_cook_tone2: 👨🏽‍🍳 :man_cook_tone3: 👨🏾‍🍳 :man_cook_tone4: 👨🏿‍🍳 :man_cook_tone5: 👩‍🍳 :woman_cook:
//👩🏻‍🍳 :woman_cook_tone1: 👩🏼‍🍳 :woman_cook_tone2: 👩🏽‍🍳 :woman_cook_tone3: 👩🏾‍🍳 :woman_cook_tone4: 👩🏿‍🍳 :woman_cook_tone5: 👨‍🔧 :man_mechanic: 👨🏻‍🔧 :man_mechanic_tone1: 👨🏼‍🔧 :man_mechanic_tone2: 👨🏽‍🔧 :man_mechanic_tone3:
//👨🏾‍🔧 :man_mechanic_tone4: 👨🏿‍🔧 :man_mechanic_tone5: 👩‍🔧 :woman_mechanic: 👩🏻‍🔧 :woman_mechanic_tone1: 👩🏼‍🔧 :woman_mechanic_tone2: 👩🏽‍🔧 :woman_mechanic_tone3: 👩🏾‍🔧 :woman_mechanic_tone4: 👩🏿‍🔧 :woman_mechanic_tone5: 👨‍🏭 :man_factory_worker:
//👨🏻‍🏭 :man_factory_worker_tone1: 👨🏼‍🏭 :man_factory_worker_tone2: 👨🏽‍🏭 :man_factory_worker_tone3: 👨🏾‍🏭 :man_factory_worker_tone4: 👨🏿‍🏭 :man_factory_worker_tone5: 👩‍🏭 :woman_factory_worker: 👩🏻‍🏭 :woman_factory_worker_tone1: 👩🏼‍🏭 :woman_factory_worker_tone2: 👩🏽‍🏭 :woman_factory_worker_tone3:
//👩🏾‍🏭 :woman_factory_worker_tone4: 👩🏿‍🏭 :woman_factory_worker_tone5: 👨‍💼 :man_office_worker: 👨🏻‍💼 :man_office_worker_tone1: 👨🏼‍💼 :man_office_worker_tone2: 👨🏽‍💼 :man_office_worker_tone3: 👨🏾‍💼 :man_office_worker_tone4: 👨🏿‍💼 :man_office_worker_tone5: 👩‍💼 :woman_office_worker:
//👩🏻‍💼 :woman_office_worker_tone1: 👩🏼‍💼 :woman_office_worker_tone2: 👩🏽‍💼 :woman_office_worker_tone3: 👩🏾‍💼 :woman_office_worker_tone4: 👩🏿‍💼 :woman_office_worker_tone5: 👨‍🔬 :man_scientist: 👨🏻‍🔬 :man_scientist_tone1: 👨🏼‍🔬 :man_scientist_tone2: 👨🏽‍🔬 :man_scientist_tone3:
//👨🏾‍🔬 :man_scientist_tone4: 👨🏿‍🔬 :man_scientist_tone5: 👩‍🔬 :woman_scientist: 👩🏻‍🔬 :woman_scientist_tone1: 👩🏼‍🔬 :woman_scientist_tone2: 👩🏽‍🔬 :woman_scientist_tone3: 👩🏾‍🔬 :woman_scientist_tone4: 👩🏿‍🔬 :woman_scientist_tone5: 👨‍💻 :man_technologist:
//👨🏻‍💻 :man_technologist_tone1: 👨🏼‍💻 :man_technologist_tone2: 👨🏽‍💻 :man_technologist_tone3: 👨🏾‍💻 :man_technologist_tone4: 👨🏿‍💻 :man_technologist_tone5: 👩‍💻 :woman_technologist: 👩🏻‍💻 :woman_technologist_tone1: 👩🏼‍💻 :woman_technologist_tone2: 👩🏽‍💻 :woman_technologist_tone3:
//👩🏾‍💻 :woman_technologist_tone4: 👩🏿‍💻 :woman_technologist_tone5: 👨‍🎤 :man_singer: 👨🏻‍🎤 :man_singer_tone1: 👨🏼‍🎤 :man_singer_tone2: 👨🏽‍🎤 :man_singer_tone3: 👨🏾‍🎤 :man_singer_tone4: 👨🏿‍🎤 :man_singer_tone5: 👩‍🎤 :woman_singer:
//👩🏻‍🎤 :woman_singer_tone1: 👩🏼‍🎤 :woman_singer_tone2: 👩🏽‍🎤 :woman_singer_tone3: 👩🏾‍🎤 :woman_singer_tone4: 👩🏿‍🎤 :woman_singer_tone5: 👨‍🎨 :man_artist: 👨🏻‍🎨 :man_artist_tone1: 👨🏼‍🎨 :man_artist_tone2: 👨🏽‍🎨 :man_artist_tone3:
//👨🏾‍🎨 :man_artist_tone4: 👨🏿‍🎨 :man_artist_tone5: 👩‍🎨 :woman_artist: 👩🏻‍🎨 :woman_artist_tone1: 👩🏼‍🎨 :woman_artist_tone2: 👩🏽‍🎨 :woman_artist_tone3: 👩🏾‍🎨 :woman_artist_tone4: 👩🏿‍🎨 :woman_artist_tone5: 👨‍✈️ :man_pilot:
//👨🏻‍✈️ :man_pilot_tone1: 👨🏼‍✈️ :man_pilot_tone2: 👨🏽‍✈️ :man_pilot_tone3: 👨🏾‍✈️ :man_pilot_tone4: 👨🏿‍✈️ :man_pilot_tone5: 👩‍✈️ :woman_pilot: 👩🏻‍✈️ :woman_pilot_tone1: 👩🏼‍✈️ :woman_pilot_tone2: 👩🏽‍✈️ :woman_pilot_tone3:
//👩🏾‍✈️ :woman_pilot_tone4: 👩🏿‍✈️ :woman_pilot_tone5: 👨‍🚀 :man_astronaut: 👨🏻‍🚀 :man_astronaut_tone1: 👨🏼‍🚀 :man_astronaut_tone2: 👨🏽‍🚀 :man_astronaut_tone3: 👨🏾‍🚀 :man_astronaut_tone4: 👨🏿‍🚀 :man_astronaut_tone5: 👩‍🚀 :woman_astronaut:
//👩🏻‍🚀 :woman_astronaut_tone1: 👩🏼‍🚀 :woman_astronaut_tone2: 👩🏽‍🚀 :woman_astronaut_tone3: 👩🏾‍🚀 :woman_astronaut_tone4: 👩🏿‍🚀 :woman_astronaut_tone5: 👨‍🚒 :man_firefighter: 👨🏻‍🚒 :man_firefighter_tone1: 👨🏼‍🚒 :man_firefighter_tone2: 👨🏽‍🚒 :man_firefighter_tone3:
//👨🏾‍🚒 :man_firefighter_tone4: 👨🏿‍🚒 :man_firefighter_tone5: 👩‍🚒 :woman_firefighter: 👩🏻‍🚒 :woman_firefighter_tone1: 👩🏼‍🚒 :woman_firefighter_tone2: 👩🏽‍🚒 :woman_firefighter_tone3: 👩🏾‍🚒 :woman_firefighter_tone4: 👩🏿‍🚒 :woman_firefighter_tone5: 👮 :police_officer:
//👮🏻 :police_officer_tone1: 👮🏼 :police_officer_tone2: 👮🏽 :police_officer_tone3: 👮🏾 :police_officer_tone4: 👮🏿 :police_officer_tone5: 👮‍♂️ :man_police_officer: 👮🏻‍♂️ :man_police_officer_tone1: 👮🏼‍♂️ :man_police_officer_tone2: 👮🏽‍♂️ :man_police_officer_tone3:
//👮🏾‍♂️ :man_police_officer_tone4: 👮🏿‍♂️ :man_police_officer_tone5: 👮‍♀️ :woman_police_officer: 👮🏻‍♀️ :woman_police_officer_tone1: 👮🏼‍♀️ :woman_police_officer_tone2: 👮🏽‍♀️ :woman_police_officer_tone3: 👮🏾‍♀️ :woman_police_officer_tone4: 👮🏿‍♀️ :woman_police_officer_tone5: 🕵 :detective:
//🕵🏻 :detective_tone1: 🕵🏼 :detective_tone2: 🕵🏽 :detective_tone3: 🕵🏾 :detective_tone4: 🕵🏿 :detective_tone5: 🕵️‍♂️ :man_detective: 🕵🏻‍♂️ :man_detective_tone1: 🕵🏼‍♂️ :man_detective_tone2: 🕵🏽‍♂️ :man_detective_tone3:
//🕵🏾‍♂️ :man_detective_tone4: 🕵🏿‍♂️ :man_detective_tone5: 🕵️‍♀️ :woman_detective: 🕵🏻‍♀️ :woman_detective_tone1: 🕵🏼‍♀️ :woman_detective_tone2: 🕵🏽‍♀️ :woman_detective_tone3: 🕵🏾‍♀️ :woman_detective_tone4: 🕵🏿‍♀️ :woman_detective_tone5: 💂 :guard:
//💂🏻 :guard_tone1: 💂🏼 :guard_tone2: 💂🏽 :guard_tone3: 💂🏾 :guard_tone4: 💂🏿 :guard_tone5: 💂‍♂️ :man_guard: 💂🏻‍♂️ :man_guard_tone1: 💂🏼‍♂️ :man_guard_tone2: 💂🏽‍♂️ :man_guard_tone3:
//💂🏾‍♂️ :man_guard_tone4: 💂🏿‍♂️ :man_guard_tone5: 💂‍♀️ :woman_guard: 💂🏻‍♀️ :woman_guard_tone1: 💂🏼‍♀️ :woman_guard_tone2: 💂🏽‍♀️ :woman_guard_tone3: 💂🏾‍♀️ :woman_guard_tone4: 💂🏿‍♀️ :woman_guard_tone5: 👷 :construction_worker:
//👷🏻 :construction_worker_tone1: 👷🏼 :construction_worker_tone2: 👷🏽 :construction_worker_tone3: 👷🏾 :construction_worker_tone4: 👷🏿 :construction_worker_tone5: 👷‍♂️ :man_construction_worker: 👷🏻‍♂️ :man_construction_worker_tone1: 👷🏼‍♂️ :man_construction_worker_tone2: 👷🏽‍♂️ :man_construction_worker_tone3:
//👷🏾‍♂️ :man_construction_worker_tone4: 👷🏿‍♂️ :man_construction_worker_tone5: 👷‍♀️ :woman_construction_worker: 👷🏻‍♀️ :woman_construction_worker_tone1: 👷🏼‍♀️ :woman_construction_worker_tone2: 👷🏽‍♀️ :woman_construction_worker_tone3: 👷🏾‍♀️ :woman_construction_worker_tone4: 👷🏿‍♀️ :woman_construction_worker_tone5: 🤴 :prince:
//🤴🏻 :prince_tone1: 🤴🏼 :prince_tone2: 🤴🏽 :prince_tone3: 🤴🏾 :prince_tone4: 🤴🏿 :prince_tone5: 👸 :princess: 👸🏻 :princess_tone1: 👸🏼 :princess_tone2: 👸🏽 :princess_tone3:
//👸🏾 :princess_tone4: 👸🏿 :princess_tone5: 👳 :person_wearing_turban: 👳🏻 :person_wearing_turban_tone1: 👳🏼 :person_wearing_turban_tone2: 👳🏽 :person_wearing_turban_tone3: 👳🏾 :person_wearing_turban_tone4: 👳🏿 :person_wearing_turban_tone5: 👳‍♂️ :man_wearing_turban:
//👳🏻‍♂️ :man_wearing_turban_tone1: 👳🏼‍♂️ :man_wearing_turban_tone2: 👳🏽‍♂️ :man_wearing_turban_tone3: 👳🏾‍♂️ :man_wearing_turban_tone4: 👳🏿‍♂️ :man_wearing_turban_tone5: 👳‍♀️ :woman_wearing_turban: 👳🏻‍♀️ :woman_wearing_turban_tone1: 👳🏼‍♀️ :woman_wearing_turban_tone2: 👳🏽‍♀️ :woman_wearing_turban_tone3:
//👳🏾‍♀️ :woman_wearing_turban_tone4: 👳🏿‍♀️ :woman_wearing_turban_tone5: 👲 :man_with_chinese_cap: 👲🏻 :man_with_chinese_cap_tone1: 👲🏼 :man_with_chinese_cap_tone2: 👲🏽 :man_with_chinese_cap_tone3: 👲🏾 :man_with_chinese_cap_tone4: 👲🏿 :man_with_chinese_cap_tone5: 🧕 :woman_with_headscarf:
//🧕🏻 :woman_with_headscarf_tone1: 🧕🏼 :woman_with_headscarf_tone2: 🧕🏽 :woman_with_headscarf_tone3: 🧕🏾 :woman_with_headscarf_tone4: 🧕🏿 :woman_with_headscarf_tone5: 🧔 :bearded_person: 🧔🏻 :bearded_person_tone1: 🧔🏼 :bearded_person_tone2: 🧔🏽 :bearded_person_tone3:
//🧔🏾 :bearded_person_tone4: 🧔🏿 :bearded_person_tone5: 👱 :blond_haired_person: 👱🏻 :blond_haired_person_tone1: 👱🏼 :blond_haired_person_tone2: 👱🏽 :blond_haired_person_tone3: 👱🏾 :blond_haired_person_tone4: 👱🏿 :blond_haired_person_tone5: 👱‍♂️ :blond-haired_man:
//👱🏻‍♂️ :blond-haired_man_tone1: 👱🏼‍♂️ :blond-haired_man_tone2: 👱🏽‍♂️ :blond-haired_man_tone3: 👱🏾‍♂️ :blond-haired_man_tone4: 👱🏿‍♂️ :blond-haired_man_tone5: 👱‍♀️ :blond-haired_woman: 👱🏻‍♀️ :blond-haired_woman_tone1: 👱🏼‍♀️ :blond-haired_woman_tone2: 👱🏽‍♀️ :blond-haired_woman_tone3:
//👱🏾‍♀️ :blond-haired_woman_tone4: 👱🏿‍♀️ :blond-haired_woman_tone5: 🤵 :man_in_tuxedo: 🤵🏻 :man_in_tuxedo_tone1: 🤵🏼 :man_in_tuxedo_tone2: 🤵🏽 :man_in_tuxedo_tone3: 🤵🏾 :man_in_tuxedo_tone4: 🤵🏿 :man_in_tuxedo_tone5: 👰 :bride_with_veil:
//👰🏻 :bride_with_veil_tone1: 👰🏼 :bride_with_veil_tone2: 👰🏽 :bride_with_veil_tone3: 👰🏾 :bride_with_veil_tone4: 👰🏿 :bride_with_veil_tone5: 🤰 :pregnant_woman: 🤰🏻 :pregnant_woman_tone1: 🤰🏼 :pregnant_woman_tone2: 🤰🏽 :pregnant_woman_tone3:
//🤰🏾 :pregnant_woman_tone4: 🤰🏿 :pregnant_woman_tone5: 🤱 :breast_feeding: 🤱🏻 :breast_feeding_tone1: 🤱🏼 :breast_feeding_tone2: 🤱🏽 :breast_feeding_tone3: 🤱🏾 :breast_feeding_tone4: 🤱🏿 :breast_feeding_tone5: 👼 :angel:
//👼🏻 :angel_tone1: 👼🏼 :angel_tone2: 👼🏽 :angel_tone3: 👼🏾 :angel_tone4: 👼🏿 :angel_tone5: 🎅 :santa: 🎅🏻 :santa_tone1: 🎅🏼 :santa_tone2: 🎅🏽 :santa_tone3:
//🎅🏾 :santa_tone4: 🎅🏿 :santa_tone5: 🤶 :mrs_claus: 🤶🏻 :mrs_claus_tone1: 🤶🏼 :mrs_claus_tone2: 🤶🏽 :mrs_claus_tone3: 🤶🏾 :mrs_claus_tone4: 🤶🏿 :mrs_claus_tone5: 🧙 :mage:
//🧙🏻 :mage_tone1: 🧙🏼 :mage_tone2: 🧙🏽 :mage_tone3: 🧙🏾 :mage_tone4: 🧙🏿 :mage_tone5: 🧙‍♀️ :woman_mage: 🧙🏻‍♀️ :woman_mage_tone1: 🧙🏼‍♀️ :woman_mage_tone2: 🧙🏽‍♀️ :woman_mage_tone3:
//🧙🏾‍♀️ :woman_mage_tone4: 🧙🏿‍♀️ :woman_mage_tone5: 🧙‍♂️ :man_mage: 🧙🏻‍♂️ :man_mage_tone1: 🧙🏼‍♂️ :man_mage_tone2: 🧙🏽‍♂️ :man_mage_tone3: 🧙🏾‍♂️ :man_mage_tone4: 🧙🏿‍♂️ :man_mage_tone5: 🧚 :fairy:
//🧚🏻 :fairy_tone1: 🧚🏼 :fairy_tone2: 🧚🏽 :fairy_tone3: 🧚🏾 :fairy_tone4: 🧚🏿 :fairy_tone5: 🧚‍♀️ :woman_fairy: 🧚🏻‍♀️ :woman_fairy_tone1: 🧚🏼‍♀️ :woman_fairy_tone2: 🧚🏽‍♀️ :woman_fairy_tone3:
//🧚🏾‍♀️ :woman_fairy_tone4: 🧚🏿‍♀️ :woman_fairy_tone5: 🧚‍♂️ :man_fairy: 🧚🏻‍♂️ :man_fairy_tone1: 🧚🏼‍♂️ :man_fairy_tone2: 🧚🏽‍♂️ :man_fairy_tone3: 🧚🏾‍♂️ :man_fairy_tone4: 🧚🏿‍♂️ :man_fairy_tone5: 🧛 :vampire:
//🧛🏻 :vampire_tone1: 🧛🏼 :vampire_tone2: 🧛🏽 :vampire_tone3: 🧛🏾 :vampire_tone4: 🧛🏿 :vampire_tone5: 🧛‍♀️ :woman_vampire: 🧛🏻‍♀️ :woman_vampire_tone1: 🧛🏼‍♀️ :woman_vampire_tone2: 🧛🏽‍♀️ :woman_vampire_tone3:
//🧛🏾‍♀️ :woman_vampire_tone4: 🧛🏿‍♀️ :woman_vampire_tone5: 🧛‍♂️ :man_vampire: 🧛🏻‍♂️ :man_vampire_tone1: 🧛🏼‍♂️ :man_vampire_tone2: 🧛🏽‍♂️ :man_vampire_tone3: 🧛🏾‍♂️ :man_vampire_tone4: 🧛🏿‍♂️ :man_vampire_tone5: 🧜 :merperson:
//🧜🏻 :merperson_tone1: 🧜🏼 :merperson_tone2: 🧜🏽 :merperson_tone3: 🧜🏾 :merperson_tone4: 🧜🏿 :merperson_tone5: 🧜‍♀️ :mermaid: 🧜🏻‍♀️ :mermaid_tone1: 🧜🏼‍♀️ :mermaid_tone2: 🧜🏽‍♀️ :mermaid_tone3:
//🧜🏾‍♀️ :mermaid_tone4: 🧜🏿‍♀️ :mermaid_tone5: 🧜‍♂️ :merman: 🧜🏻‍♂️ :merman_tone1: 🧜🏼‍♂️ :merman_tone2: 🧜🏽‍♂️ :merman_tone3: 🧜🏾‍♂️ :merman_tone4: 🧜🏿‍♂️ :merman_tone5: 🧝 :elf:
//🧝🏻 :elf_tone1: 🧝🏼 :elf_tone2: 🧝🏽 :elf_tone3: 🧝🏾 :elf_tone4: 🧝🏿 :elf_tone5: 🧝‍♀️ :woman_elf: 🧝🏻‍♀️ :woman_elf_tone1: 🧝🏼‍♀️ :woman_elf_tone2: 🧝🏽‍♀️ :woman_elf_tone3:
//🧝🏾‍♀️ :woman_elf_tone4: 🧝🏿‍♀️ :woman_elf_tone5: 🧝‍♂️ :man_elf: 🧝🏻‍♂️ :man_elf_tone1: 🧝🏼‍♂️ :man_elf_tone2: 🧝🏽‍♂️ :man_elf_tone3: 🧝🏾‍♂️ :man_elf_tone4: 🧝🏿‍♂️ :man_elf_tone5: 🧞 :genie:
//🧞‍♀️ :woman_genie: 🧞‍♂️ :man_genie: 🧟 :zombie: 🧟‍♀️ :woman_zombie: 🧟‍♂️ :man_zombie: 🙍 :person_frowning: 🙍🏻 :person_frowning_tone1: 🙍🏼 :person_frowning_tone2: 🙍🏽 :person_frowning_tone3:
//🙍🏾 :person_frowning_tone4: 🙍🏿 :person_frowning_tone5: 🙍‍♂️ :man_frowning: 🙍🏻‍♂️ :man_frowning_tone1: 🙍🏼‍♂️ :man_frowning_tone2: 🙍🏽‍♂️ :man_frowning_tone3: 🙍🏾‍♂️ :man_frowning_tone4: 🙍🏿‍♂️ :man_frowning_tone5: 🙍‍♀️ :woman_frowning:
//🙍🏻‍♀️ :woman_frowning_tone1: 🙍🏼‍♀️ :woman_frowning_tone2: 🙍🏽‍♀️ :woman_frowning_tone3: 🙍🏾‍♀️ :woman_frowning_tone4: 🙍🏿‍♀️ :woman_frowning_tone5: 🙎 :person_pouting: 🙎🏻 :person_pouting_tone1: 🙎🏼 :person_pouting_tone2: 🙎🏽 :person_pouting_tone3:
//🙎🏾 :person_pouting_tone4: 🙎🏿 :person_pouting_tone5: 🙎‍♂️ :man_pouting: 🙎🏻‍♂️ :man_pouting_tone1: 🙎🏼‍♂️ :man_pouting_tone2: 🙎🏽‍♂️ :man_pouting_tone3: 🙎🏾‍♂️ :man_pouting_tone4: 🙎🏿‍♂️ :man_pouting_tone5: 🙎‍♀️ :woman_pouting:
//🙎🏻‍♀️ :woman_pouting_tone1: 🙎🏼‍♀️ :woman_pouting_tone2: 🙎🏽‍♀️ :woman_pouting_tone3: 🙎🏾‍♀️ :woman_pouting_tone4: 🙎🏿‍♀️ :woman_pouting_tone5: 🙅 :person_gesturing_no: 🙅🏻 :person_gesturing_no_tone1: 🙅🏼 :person_gesturing_no_tone2: 🙅🏽 :person_gesturing_no_tone3:
//🙅🏾 :person_gesturing_no_tone4: 🙅🏿 :person_gesturing_no_tone5: 🙅‍♂️ :man_gesturing_no: 🙅🏻‍♂️ :man_gesturing_no_tone1: 🙅🏼‍♂️ :man_gesturing_no_tone2: 🙅🏽‍♂️ :man_gesturing_no_tone3: 🙅🏾‍♂️ :man_gesturing_no_tone4: 🙅🏿‍♂️ :man_gesturing_no_tone5: 🙅‍♀️ :woman_gesturing_no:
//🙅🏻‍♀️ :woman_gesturing_no_tone1: 🙅🏼‍♀️ :woman_gesturing_no_tone2: 🙅🏽‍♀️ :woman_gesturing_no_tone3: 🙅🏾‍♀️ :woman_gesturing_no_tone4: 🙅🏿‍♀️ :woman_gesturing_no_tone5: 🙆 :person_gesturing_ok: 🙆🏻 :person_gesturing_ok_tone1: 🙆🏼 :person_gesturing_ok_tone2: 🙆🏽 :person_gesturing_ok_tone3:
//🙆🏾 :person_gesturing_ok_tone4: 🙆🏿 :person_gesturing_ok_tone5: 🙆‍♂️ :man_gesturing_ok: 🙆🏻‍♂️ :man_gesturing_ok_tone1: 🙆🏼‍♂️ :man_gesturing_ok_tone2: 🙆🏽‍♂️ :man_gesturing_ok_tone3: 🙆🏾‍♂️ :man_gesturing_ok_tone4: 🙆🏿‍♂️ :man_gesturing_ok_tone5: 🙆‍♀️ :woman_gesturing_ok:
//🙆🏻‍♀️ :woman_gesturing_ok_tone1: 🙆🏼‍♀️ :woman_gesturing_ok_tone2: 🙆🏽‍♀️ :woman_gesturing_ok_tone3: 🙆🏾‍♀️ :woman_gesturing_ok_tone4: 🙆🏿‍♀️ :woman_gesturing_ok_tone5: 💁 :person_tipping_hand: 💁🏻 :person_tipping_hand_tone1: 💁🏼 :person_tipping_hand_tone2: 💁🏽 :person_tipping_hand_tone3:
//💁🏾 :person_tipping_hand_tone4: 💁🏿 :person_tipping_hand_tone5: 💁‍♂️ :man_tipping_hand: 💁🏻‍♂️ :man_tipping_hand_tone1: 💁🏼‍♂️ :man_tipping_hand_tone2: 💁🏽‍♂️ :man_tipping_hand_tone3: 💁🏾‍♂️ :man_tipping_hand_tone4: 💁🏿‍♂️ :man_tipping_hand_tone5: 💁‍♀️ :woman_tipping_hand:
//💁🏻‍♀️ :woman_tipping_hand_tone1: 💁🏼‍♀️ :woman_tipping_hand_tone2: 💁🏽‍♀️ :woman_tipping_hand_tone3: 💁🏾‍♀️ :woman_tipping_hand_tone4: 💁🏿‍♀️ :woman_tipping_hand_tone5: 🙋 :person_raising_hand: 🙋🏻 :person_raising_hand_tone1: 🙋🏼 :person_raising_hand_tone2: 🙋🏽 :person_raising_hand_tone3:
//🙋🏾 :person_raising_hand_tone4: 🙋🏿 :person_raising_hand_tone5: 🙋‍♂️ :man_raising_hand: 🙋🏻‍♂️ :man_raising_hand_tone1: 🙋🏼‍♂️ :man_raising_hand_tone2: 🙋🏽‍♂️ :man_raising_hand_tone3: 🙋🏾‍♂️ :man_raising_hand_tone4: 🙋🏿‍♂️ :man_raising_hand_tone5: 🙋‍♀️ :woman_raising_hand:
//🙋🏻‍♀️ :woman_raising_hand_tone1: 🙋🏼‍♀️ :woman_raising_hand_tone2: 🙋🏽‍♀️ :woman_raising_hand_tone3: 🙋🏾‍♀️ :woman_raising_hand_tone4: 🙋🏿‍♀️ :woman_raising_hand_tone5: 🙇 :person_bowing: 🙇🏻 :person_bowing_tone1: 🙇🏼 :person_bowing_tone2: 🙇🏽 :person_bowing_tone3:
//🙇🏾 :person_bowing_tone4: 🙇🏿 :person_bowing_tone5: 🙇‍♂️ :man_bowing: 🙇🏻‍♂️ :man_bowing_tone1: 🙇🏼‍♂️ :man_bowing_tone2: 🙇🏽‍♂️ :man_bowing_tone3: 🙇🏾‍♂️ :man_bowing_tone4: 🙇🏿‍♂️ :man_bowing_tone5: 🙇‍♀️ :woman_bowing:
//🙇🏻‍♀️ :woman_bowing_tone1: 🙇🏼‍♀️ :woman_bowing_tone2: 🙇🏽‍♀️ :woman_bowing_tone3: 🙇🏾‍♀️ :woman_bowing_tone4: 🙇🏿‍♀️ :woman_bowing_tone5: 🤦 :person_facepalming: 🤦🏻 :person_facepalming_tone1: 🤦🏼 :person_facepalming_tone2: 🤦🏽 :person_facepalming_tone3:
//🤦🏾 :person_facepalming_tone4: 🤦🏿 :person_facepalming_tone5: 🤦‍♂️ :man_facepalming: 🤦🏻‍♂️ :man_facepalming_tone1: 🤦🏼‍♂️ :man_facepalming_tone2: 🤦🏽‍♂️ :man_facepalming_tone3: 🤦🏾‍♂️ :man_facepalming_tone4: 🤦🏿‍♂️ :man_facepalming_tone5: 🤦‍♀️ :woman_facepalming:
//🤦🏻‍♀️ :woman_facepalming_tone1: 🤦🏼‍♀️ :woman_facepalming_tone2: 🤦🏽‍♀️ :woman_facepalming_tone3: 🤦🏾‍♀️ :woman_facepalming_tone4: 🤦🏿‍♀️ :woman_facepalming_tone5: 🤷 :person_shrugging: 🤷🏻 :person_shrugging_tone1: 🤷🏼 :person_shrugging_tone2: 🤷🏽 :person_shrugging_tone3:
//🤷🏾 :person_shrugging_tone4: 🤷🏿 :person_shrugging_tone5: 🤷‍♂️ :man_shrugging: 🤷🏻‍♂️ :man_shrugging_tone1: 🤷🏼‍♂️ :man_shrugging_tone2: 🤷🏽‍♂️ :man_shrugging_tone3: 🤷🏾‍♂️ :man_shrugging_tone4: 🤷🏿‍♂️ :man_shrugging_tone5: 🤷‍♀️ :woman_shrugging:
//🤷🏻‍♀️ :woman_shrugging_tone1: 🤷🏼‍♀️ :woman_shrugging_tone2: 🤷🏽‍♀️ :woman_shrugging_tone3: 🤷🏾‍♀️ :woman_shrugging_tone4: 🤷🏿‍♀️ :woman_shrugging_tone5: 💆 :person_getting_massage: 💆🏻 :person_getting_massage_tone1: 💆🏼 :person_getting_massage_tone2: 💆🏽 :person_getting_massage_tone3:
//💆🏾 :person_getting_massage_tone4: 💆🏿 :person_getting_massage_tone5: 💆‍♂️ :man_getting_face_massage: 💆🏻‍♂️ :man_getting_face_massage_tone1: 💆🏼‍♂️ :man_getting_face_massage_tone2: 💆🏽‍♂️ :man_getting_face_massage_tone3: 💆🏾‍♂️ :man_getting_face_massage_tone4: 💆🏿‍♂️ :man_getting_face_massage_tone5: 💆‍♀️ :woman_getting_face_massage:
//💆🏻‍♀️ :woman_getting_face_massage_tone1: 💆🏼‍♀️ :woman_getting_face_massage_tone2: 💆🏽‍♀️ :woman_getting_face_massage_tone3: 💆🏾‍♀️ :woman_getting_face_massage_tone4: 💆🏿‍♀️ :woman_getting_face_massage_tone5: 💇 :person_getting_haircut: 💇🏻 :person_getting_haircut_tone1: 💇🏼 :person_getting_haircut_tone2: 💇🏽 :person_getting_haircut_tone3:
//💇🏾 :person_getting_haircut_tone4: 💇🏿 :person_getting_haircut_tone5: 💇‍♂️ :man_getting_haircut: 💇🏻‍♂️ :man_getting_haircut_tone1: 💇🏼‍♂️ :man_getting_haircut_tone2: 💇🏽‍♂️ :man_getting_haircut_tone3: 💇🏾‍♂️ :man_getting_haircut_tone4: 💇🏿‍♂️ :man_getting_haircut_tone5: 💇‍♀️ :woman_getting_haircut:
//💇🏻‍♀️ :woman_getting_haircut_tone1: 💇🏼‍♀️ :woman_getting_haircut_tone2: 💇🏽‍♀️ :woman_getting_haircut_tone3: 💇🏾‍♀️ :woman_getting_haircut_tone4: 💇🏿‍♀️ :woman_getting_haircut_tone5: 🚶 :person_walking: 🚶🏻 :person_walking_tone1: 🚶🏼 :person_walking_tone2: 🚶🏽 :person_walking_tone3:
//🚶🏾 :person_walking_tone4: 🚶🏿 :person_walking_tone5: 🚶‍♂️ :man_walking: 🚶🏻‍♂️ :man_walking_tone1: 🚶🏼‍♂️ :man_walking_tone2: 🚶🏽‍♂️ :man_walking_tone3: 🚶🏾‍♂️ :man_walking_tone4: 🚶🏿‍♂️ :man_walking_tone5: 🚶‍♀️ :woman_walking:
//🚶🏻‍♀️ :woman_walking_tone1: 🚶🏼‍♀️ :woman_walking_tone2: 🚶🏽‍♀️ :woman_walking_tone3: 🚶🏾‍♀️ :woman_walking_tone4: 🚶🏿‍♀️ :woman_walking_tone5: 🏃 :person_running: 🏃🏻 :person_running_tone1: 🏃🏼 :person_running_tone2: 🏃🏽 :person_running_tone3:
//🏃🏾 :person_running_tone4: 🏃🏿 :person_running_tone5: 🏃‍♂️ :man_running: 🏃🏻‍♂️ :man_running_tone1: 🏃🏼‍♂️ :man_running_tone2: 🏃🏽‍♂️ :man_running_tone3: 🏃🏾‍♂️ :man_running_tone4: 🏃🏿‍♂️ :man_running_tone5: 🏃‍♀️ :woman_running:
//🏃🏻‍♀️ :woman_running_tone1: 🏃🏼‍♀️ :woman_running_tone2: 🏃🏽‍♀️ :woman_running_tone3: 🏃🏾‍♀️ :woman_running_tone4: 🏃🏿‍♀️ :woman_running_tone5: 💃 :dancer: 💃🏻 :dancer_tone1: 💃🏼 :dancer_tone2: 💃🏽 :dancer_tone3:
//💃🏾 :dancer_tone4: 💃🏿 :dancer_tone5: 🕺 :man_dancing: 🕺🏻 :man_dancing_tone1: 🕺🏼 :man_dancing_tone2: 🕺🏽 :man_dancing_tone3: 🕺🏾 :man_dancing_tone4: 🕺🏿 :man_dancing_tone5: 👯 :people_with_bunny_ears_partying:
//👯‍♂️ :men_with_bunny_ears_partying: 👯‍♀️ :women_with_bunny_ears_partying: 🧖 :person_in_steamy_room: 🧖🏻 :person_in_steamy_room_tone1: 🧖🏼 :person_in_steamy_room_tone2: 🧖🏽 :person_in_steamy_room_tone3: 🧖🏾 :person_in_steamy_room_tone4: 🧖🏿 :person_in_steamy_room_tone5: 🧖‍♀️ :woman_in_steamy_room:
//🧖🏻‍♀️ :woman_in_steamy_room_tone1: 🧖🏼‍♀️ :woman_in_steamy_room_tone2: 🧖🏽‍♀️ :woman_in_steamy_room_tone3: 🧖🏾‍♀️ :woman_in_steamy_room_tone4: 🧖🏿‍♀️ :woman_in_steamy_room_tone5: 🧖‍♂️ :man_in_steamy_room: 🧖🏻‍♂️ :man_in_steamy_room_tone1: 🧖🏼‍♂️ :man_in_steamy_room_tone2: 🧖🏽‍♂️ :man_in_steamy_room_tone3:
//🧖🏾‍♂️ :man_in_steamy_room_tone4: 🧖🏿‍♂️ :man_in_steamy_room_tone5: 🧗 :person_climbing: 🧗🏻 :person_climbing_tone1: 🧗🏼 :person_climbing_tone2: 🧗🏽 :person_climbing_tone3: 🧗🏾 :person_climbing_tone4: 🧗🏿 :person_climbing_tone5: 🧗‍♀️ :woman_climbing:
//🧗🏻‍♀️ :woman_climbing_tone1: 🧗🏼‍♀️ :woman_climbing_tone2: 🧗🏽‍♀️ :woman_climbing_tone3: 🧗🏾‍♀️ :woman_climbing_tone4: 🧗🏿‍♀️ :woman_climbing_tone5: 🧗‍♂️ :man_climbing: 🧗🏻‍♂️ :man_climbing_tone1: 🧗🏼‍♂️ :man_climbing_tone2: 🧗🏽‍♂️ :man_climbing_tone3:
//🧗🏾‍♂️ :man_climbing_tone4: 🧗🏿‍♂️ :man_climbing_tone5: 🧘 :person_in_lotus_position: 🧘🏻 :person_in_lotus_position_tone1: 🧘🏼 :person_in_lotus_position_tone2: 🧘🏽 :person_in_lotus_position_tone3: 🧘🏾 :person_in_lotus_position_tone4: 🧘🏿 :person_in_lotus_position_tone5: 🧘‍♀️ :woman_in_lotus_position:
//🧘🏻‍♀️ :woman_in_lotus_position_tone1: 🧘🏼‍♀️ :woman_in_lotus_position_tone2: 🧘🏽‍♀️ :woman_in_lotus_position_tone3: 🧘🏾‍♀️ :woman_in_lotus_position_tone4: 🧘🏿‍♀️ :woman_in_lotus_position_tone5: 🧘‍♂️ :man_in_lotus_position: 🧘🏻‍♂️ :man_in_lotus_position_tone1: 🧘🏼‍♂️ :man_in_lotus_position_tone2: 🧘🏽‍♂️ :man_in_lotus_position_tone3:
//🧘🏾‍♂️ :man_in_lotus_position_tone4: 🧘🏿‍♂️ :man_in_lotus_position_tone5: 🛀 :bath: 🛀🏻 :bath_tone1: 🛀🏼 :bath_tone2: 🛀🏽 :bath_tone3: 🛀🏾 :bath_tone4: 🛀🏿 :bath_tone5: 🛌 :sleeping_accommodation:
//🛌🏻 :person_in_bed_tone1: 🛌🏼 :person_in_bed_tone2: 🛌🏽 :person_in_bed_tone3: 🛌🏾 :person_in_bed_tone4: 🛌🏿 :person_in_bed_tone5: 🕴  🕴🏻 :man_in_business_suit_levitating_tone1: 🕴🏼 :man_in_business_suit_levitating_tone2: 🕴🏽 :man_in_business_suit_levitating_tone3:
//🕴🏾 :man_in_business_suit_levitating_tone4: 🕴🏿 :man_in_business_suit_levitating_tone5: 🗣 :speaking_head: 👤 :bust_in_silhouette: 👥 :busts_in_silhouette: 🤺 :person_fencing: 🏇 :horse_racing: 🏇🏻 :horse_racing_tone1: 🏇🏼 :horse_racing_tone2:
//🏇🏽 :horse_racing_tone3: 🏇🏾 :horse_racing_tone4: 🏇🏿 :horse_racing_tone5: ⛷ :skier: 🏂 :snowboarder: 🏂🏻 :snowboarder_tone1: 🏂🏼 :snowboarder_tone2: 🏂🏽 :snowboarder_tone3: 🏂🏾 :snowboarder_tone4:
//🏂🏿 :snowboarder_tone5: 🏌 :person_golfing: 🏌🏻 :person_golfing_tone1: 🏌🏼 :person_golfing_tone2: 🏌🏽 :person_golfing_tone3: 🏌🏾 :person_golfing_tone4: 🏌🏿 :person_golfing_tone5: 🏌️‍♂️ :man_golfing: 🏌🏻‍♂️ :man_golfing_tone1:
//🏌🏼‍♂️ :man_golfing_tone2: 🏌🏽‍♂️ :man_golfing_tone3: 🏌🏾‍♂️ :man_golfing_tone4: 🏌🏿‍♂️ :man_golfing_tone5: 🏌️‍♀️ :woman_golfing: 🏌🏻‍♀️ :woman_golfing_tone1: 🏌🏼‍♀️ :woman_golfing_tone2: 🏌🏽‍♀️ :woman_golfing_tone3: 🏌🏾‍♀️ :woman_golfing_tone4:
//🏌🏿‍♀️ :woman_golfing_tone5: 🏄 :person_surfing: 🏄🏻 :person_surfing_tone1: 🏄🏼 :person_surfing_tone2: 🏄🏽 :person_surfing_tone3: 🏄🏾 :person_surfing_tone4: 🏄🏿 :person_surfing_tone5: 🏄‍♂️ :man_surfing: 🏄🏻‍♂️ :man_surfing_tone1:
//🏄🏼‍♂️ :man_surfing_tone2: 🏄🏽‍♂️ :man_surfing_tone3: 🏄🏾‍♂️ :man_surfing_tone4: 🏄🏿‍♂️ :man_surfing_tone5: 🏄‍♀️ :woman_surfing: 🏄🏻‍♀️ :woman_surfing_tone1: 🏄🏼‍♀️ :woman_surfing_tone2: 🏄🏽‍♀️ :woman_surfing_tone3: 🏄🏾‍♀️ :woman_surfing_tone4:
//🏄🏿‍♀️ :woman_surfing_tone5: 🚣 :person_rowing_boat: 🚣🏻 :person_rowing_boat_tone1: 🚣🏼 :person_rowing_boat_tone2: 🚣🏽 :person_rowing_boat_tone3: 🚣🏾 :person_rowing_boat_tone4: 🚣🏿 :person_rowing_boat_tone5: 🚣‍♂️ :man_rowing_boat: 🚣🏻‍♂️ :man_rowing_boat_tone1:
//🚣🏼‍♂️ :man_rowing_boat_tone2: 🚣🏽‍♂️ :man_rowing_boat_tone3: 🚣🏾‍♂️ :man_rowing_boat_tone4: 🚣🏿‍♂️ :man_rowing_boat_tone5: 🚣‍♀️ :woman_rowing_boat: 🚣🏻‍♀️ :woman_rowing_boat_tone1: 🚣🏼‍♀️ :woman_rowing_boat_tone2: 🚣🏽‍♀️ :woman_rowing_boat_tone3: 🚣🏾‍♀️ :woman_rowing_boat_tone4:
//🚣🏿‍♀️ :woman_rowing_boat_tone5: 🏊 :person_swimming: 🏊🏻 :person_swimming_tone1: 🏊🏼 :person_swimming_tone2: 🏊🏽 :person_swimming_tone3: 🏊🏾 :person_swimming_tone4: 🏊🏿 :person_swimming_tone5: 🏊‍♂️ :man_swimming: 🏊🏻‍♂️ :man_swimming_tone1:
//🏊🏼‍♂️ :man_swimming_tone2: 🏊🏽‍♂️ :man_swimming_tone3: 🏊🏾‍♂️ :man_swimming_tone4: 🏊🏿‍♂️ :man_swimming_tone5: 🏊‍♀️ :woman_swimming: 🏊🏻‍♀️ :woman_swimming_tone1: 🏊🏼‍♀️ :woman_swimming_tone2: 🏊🏽‍♀️ :woman_swimming_tone3: 🏊🏾‍♀️ :woman_swimming_tone4:
//🏊🏿‍♀️ :woman_swimming_tone5: ⛹ :person_bouncing_ball: ⛹🏻 :person_bouncing_ball_tone1: ⛹🏼 :person_bouncing_ball_tone2: ⛹🏽 :person_bouncing_ball_tone3: ⛹🏾 :person_bouncing_ball_tone4: ⛹🏿 :person_bouncing_ball_tone5: ⛹️‍♂️ :man_bouncing_ball: ⛹🏻‍♂️ :man_bouncing_ball_tone1:
//⛹🏼‍♂️ :man_bouncing_ball_tone2: ⛹🏽‍♂️ :man_bouncing_ball_tone3: ⛹🏾‍♂️ :man_bouncing_ball_tone4: ⛹🏿‍♂️ :man_bouncing_ball_tone5: ⛹️‍♀️ :woman_bouncing_ball: ⛹🏻‍♀️ :woman_bouncing_ball_tone1: ⛹🏼‍♀️ :woman_bouncing_ball_tone2: ⛹🏽‍♀️ :woman_bouncing_ball_tone3: ⛹🏾‍♀️ :woman_bouncing_ball_tone4:
//⛹🏿‍♀️ :woman_bouncing_ball_tone5: 🏋 :person_lifting_weights: 🏋🏻 :person_lifting_weights_tone1: 🏋🏼 :person_lifting_weights_tone2: 🏋🏽 :person_lifting_weights_tone3: 🏋🏾 :person_lifting_weights_tone4: 🏋🏿 :person_lifting_weights_tone5: 🏋️‍♂️ :man_lifting_weights: 🏋🏻‍♂️ :man_lifting_weights_tone1:
//🏋🏼‍♂️ :man_lifting_weights_tone2: 🏋🏽‍♂️ :man_lifting_weights_tone3: 🏋🏾‍♂️ :man_lifting_weights_tone4: 🏋🏿‍♂️ :man_lifting_weights_tone5: 🏋️‍♀️ :woman_lifting_weights: 🏋🏻‍♀️ :woman_lifting_weights_tone1: 🏋🏼‍♀️ :woman_lifting_weights_tone2: 🏋🏽‍♀️ :woman_lifting_weights_tone3: 🏋🏾‍♀️ :woman_lifting_weights_tone4:
//🏋🏿‍♀️ :woman_lifting_weights_tone5: 🚴 :person_biking: 🚴🏻 :person_biking_tone1: 🚴🏼 :person_biking_tone2: 🚴🏽 :person_biking_tone3: 🚴🏾 :person_biking_tone4: 🚴🏿 :person_biking_tone5: 🚴‍♂️ :man_biking: 🚴🏻‍♂️ :man_biking_tone1:
//🚴🏼‍♂️ :man_biking_tone2: 🚴🏽‍♂️ :man_biking_tone3: 🚴🏾‍♂️ :man_biking_tone4: 🚴🏿‍♂️ :man_biking_tone5: 🚴‍♀️ :woman_biking: 🚴🏻‍♀️ :woman_biking_tone1: 🚴🏼‍♀️ :woman_biking_tone2: 🚴🏽‍♀️ :woman_biking_tone3: 🚴🏾‍♀️ :woman_biking_tone4:
//🚴🏿‍♀️ :woman_biking_tone5: 🚵 :person_mountain_biking: 🚵🏻 :person_mountain_biking_tone1: 🚵🏼 :person_mountain_biking_tone2: 🚵🏽 :person_mountain_biking_tone3: 🚵🏾 :person_mountain_biking_tone4: 🚵🏿 :person_mountain_biking_tone5: 🚵‍♂️ :man_mountain_biking: 🚵🏻‍♂️ :man_mountain_biking_tone1:
//🚵🏼‍♂️ :man_mountain_biking_tone2: 🚵🏽‍♂️ :man_mountain_biking_tone3: 🚵🏾‍♂️ :man_mountain_biking_tone4: 🚵🏿‍♂️ :man_mountain_biking_tone5: 🚵‍♀️ :woman_mountain_biking: 🚵🏻‍♀️ :woman_mountain_biking_tone1: 🚵🏼‍♀️ :woman_mountain_biking_tone2: 🚵🏽‍♀️ :woman_mountain_biking_tone3: 🚵🏾‍♀️ :woman_mountain_biking_tone4:
//🚵🏿‍♀️ :woman_mountain_biking_tone5: 🏎 :race_car: 🏍 :motorcycle: 🤸 :person_doing_cartwheel: 🤸🏻 :person_doing_cartwheel_tone1: 🤸🏼 :person_doing_cartwheel_tone2: 🤸🏽 :person_doing_cartwheel_tone3: 🤸🏾 :person_doing_cartwheel_tone4: 🤸🏿 :person_doing_cartwheel_tone5:
//🤸‍♂️ :man_cartwheeling: 🤸🏻‍♂️ :man_cartwheeling_tone1: 🤸🏼‍♂️ :man_cartwheeling_tone2: 🤸🏽‍♂️ :man_cartwheeling_tone3: 🤸🏾‍♂️ :man_cartwheeling_tone4: 🤸🏿‍♂️ :man_cartwheeling_tone5: 🤸‍♀️ :woman_cartwheeling: 🤸🏻‍♀️ :woman_cartwheeling_tone1: 🤸🏼‍♀️ :woman_cartwheeling_tone2:
//🤸🏽‍♀️ :woman_cartwheeling_tone3: 🤸🏾‍♀️ :woman_cartwheeling_tone4: 🤸🏿‍♀️ :woman_cartwheeling_tone5: 🤼 :people_wrestling: 🤼‍♂️ :men_wrestling: 🤼‍♀️ :women_wrestling: 🤽 :person_playing_water_polo: 🤽🏻 :person_playing_water_polo_tone1: 🤽🏼 :person_playing_water_polo_tone2:
//🤽🏽 :person_playing_water_polo_tone3: 🤽🏾 :person_playing_water_polo_tone4: 🤽🏿 :person_playing_water_polo_tone5: 🤽‍♂️ :man_playing_water_polo: 🤽🏻‍♂️ :man_playing_water_polo_tone1: 🤽🏼‍♂️ :man_playing_water_polo_tone2: 🤽🏽‍♂️ :man_playing_water_polo_tone3: 🤽🏾‍♂️ :man_playing_water_polo_tone4: 🤽🏿‍♂️ :man_playing_water_polo_tone5:
//🤽‍♀️ :woman_playing_water_polo: 🤽🏻‍♀️ :woman_playing_water_polo_tone1: 🤽🏼‍♀️ :woman_playing_water_polo_tone2: 🤽🏽‍♀️ :woman_playing_water_polo_tone3: 🤽🏾‍♀️ :woman_playing_water_polo_tone4: 🤽🏿‍♀️ :woman_playing_water_polo_tone5: 🤾 :person_playing_handball: 🤾🏻 :person_playing_handball_tone1: 🤾🏼 :person_playing_handball_tone2:
//🤾🏽 :person_playing_handball_tone3: 🤾🏾 :person_playing_handball_tone4: 🤾🏿 :person_playing_handball_tone5: 🤾‍♂️ :man_playing_handball: 🤾🏻‍♂️ :man_playing_handball_tone1: 🤾🏼‍♂️ :man_playing_handball_tone2: 🤾🏽‍♂️ :man_playing_handball_tone3: 🤾🏾‍♂️ :man_playing_handball_tone4: 🤾🏿‍♂️ :man_playing_handball_tone5:
//🤾‍♀️ :woman_playing_handball: 🤾🏻‍♀️ :woman_playing_handball_tone1: 🤾🏼‍♀️ :woman_playing_handball_tone2: 🤾🏽‍♀️ :woman_playing_handball_tone3: 🤾🏾‍♀️ :woman_playing_handball_tone4: 🤾🏿‍♀️ :woman_playing_handball_tone5: 🤹 :person_juggling: 🤹🏻 :person_juggling_tone1: 🤹🏼 :person_juggling_tone2:
//🤹🏽 :person_juggling_tone3: 🤹🏾 :person_juggling_tone4: 🤹🏿 :person_juggling_tone5: 🤹‍♂️ :man_juggling: 🤹🏻‍♂️ :man_juggling_tone1: 🤹🏼‍♂️ :man_juggling_tone2: 🤹🏽‍♂️ :man_juggling_tone3: 🤹🏾‍♂️ :man_juggling_tone4: 🤹🏿‍♂️ :man_juggling_tone5:
//🤹‍♀️ :woman_juggling: 🤹🏻‍♀️ :woman_juggling_tone1: 🤹🏼‍♀️ :woman_juggling_tone2: 🤹🏽‍♀️ :woman_juggling_tone3: 🤹🏾‍♀️ :woman_juggling_tone4: 🤹🏿‍♀️ :woman_juggling_tone5: 👫 :couple: 👬 :two_men_holding_hands: 👭 :two_women_holding_hands:
//💏 :couplekiss: 👩‍❤️‍💋‍👨 :kiss_woman_man: 👨‍❤️‍💋‍👨 :kiss_mm: 👩‍❤️‍💋‍👩 :kiss_ww: 💑 :couple_with_heart: 👩‍❤️‍👨 :couple_with_heart_woman_man: 👨‍❤️‍👨 :couple_mm: 👩‍❤️‍👩 :couple_ww: 👪 :family:
//👨‍👩‍👦 :family_man_woman_boy: 👨‍👩‍👧 :family_mwg: 👨‍👩‍👧‍👦 :family_mwgb: 👨‍👩‍👦‍👦 :family_mwbb: 👨‍👩‍👧‍👧 :family_mwgg: 👨‍👨‍👦 :family_mmb: 👨‍👨‍👧 :family_mmg: 👨‍👨‍👧‍👦 :family_mmgb: 👨‍👨‍👦‍👦 :family_mmbb:
//👨‍👨‍👧‍👧 :family_mmgg: 👩‍👩‍👦 :family_wwb: 👩‍👩‍👧 :family_wwg: 👩‍👩‍👧‍👦 :family_wwgb: 👩‍👩‍👦‍👦 :family_wwbb: 👩‍👩‍👧‍👧 :family_wwgg: 👨‍👦 :family_man_boy: 👨‍👦‍👦 :family_man_boy_boy: 👨‍👧 :family_man_girl:
//👨‍👧‍👦 :family_man_girl_boy: 👨‍👧‍👧 :family_man_girl_girl: 👩‍👦 :family_woman_boy: 👩‍👦‍👦 :family_woman_boy_boy: 👩‍👧 :family_woman_girl: 👩‍👧‍👦 :family_woman_girl_boy: 👩‍👧‍👧 :family_woman_girl_girl: 🤳 :selfie: 🤳🏻 :selfie_tone1:
//🤳🏼 :selfie_tone2: 🤳🏽 :selfie_tone3: 🤳🏾 :selfie_tone4: 🤳🏿 :selfie_tone5: 💪 :muscle: 💪🏻 :muscle_tone1: 💪🏼 :muscle_tone2: 💪🏽 :muscle_tone3: 💪🏾 :muscle_tone4:
//💪🏿 :muscle_tone5: 👈 :point_left: 👈🏻 :point_left_tone1: 👈🏼 :point_left_tone2: 👈🏽 :point_left_tone3: 👈🏾 :point_left_tone4: 👈🏿 :point_left_tone5: 👉 :point_right: 👉🏻 :point_right_tone1:
//👉🏼 :point_right_tone2: 👉🏽 :point_right_tone3: 👉🏾 :point_right_tone4: 👉🏿 :point_right_tone5: ☝ :point_up: ☝🏻 :point_up_tone1: ☝🏼 :point_up_tone2: ☝🏽 :point_up_tone3: ☝🏾 :point_up_tone4:
//☝🏿 :point_up_tone5: 👆 :point_up_2: 👆🏻 :point_up_2_tone1: 👆🏼 :point_up_2_tone2: 👆🏽 :point_up_2_tone3: 👆🏾 :point_up_2_tone4: 👆🏿 :point_up_2_tone5: 🖕 :middle_finger: 🖕🏻 :middle_finger_tone1:
//🖕🏼 :middle_finger_tone2: 🖕🏽 :middle_finger_tone3: 🖕🏾 :middle_finger_tone4: 🖕🏿 :middle_finger_tone5: 👇 :point_down: 👇🏻 :point_down_tone1: 👇🏼 :point_down_tone2: 👇🏽 :point_down_tone3: 👇🏾 :point_down_tone4:
//👇🏿 :point_down_tone5: ✌ :v: ✌🏻 :v_tone1: ✌🏼 :v_tone2: ✌🏽 :v_tone3: ✌🏾 :v_tone4: ✌🏿 :v_tone5: 🤞 :fingers_crossed: 🤞🏻 :fingers_crossed_tone1:
//🤞🏼 :fingers_crossed_tone2: 🤞🏽 :fingers_crossed_tone3: 🤞🏾 :fingers_crossed_tone4: 🤞🏿 :fingers_crossed_tone5: 🖖 :vulcan: 🖖🏻 :vulcan_tone1: 🖖🏼 :vulcan_tone2: 🖖🏽 :vulcan_tone3: 🖖🏾 :vulcan_tone4:
//🖖🏿 :vulcan_tone5: 🤘 :metal: 🤘🏻 :metal_tone1: 🤘🏼 :metal_tone2: 🤘🏽 :metal_tone3: 🤘🏾 :metal_tone4: 🤘🏿 :metal_tone5: 🤙 :call_me: 🤙🏻 :call_me_tone1:
//🤙🏼 :call_me_tone2: 🤙🏽 :call_me_tone3: 🤙🏾 :call_me_tone4: 🤙🏿 :call_me_tone5: 🖐  🖐🏻 :hand_splayed_tone1: 🖐🏼 :hand_splayed_tone2: 🖐🏽 :hand_splayed_tone3: 🖐🏾 :hand_splayed_tone4:
//🖐🏿 :hand_splayed_tone5: ✋ :raised_hand: ✋🏻 :raised_hand_tone1: ✋🏼 :raised_hand_tone2: ✋🏽 :raised_hand_tone3: ✋🏾 :raised_hand_tone4: ✋🏿 :raised_hand_tone5: 👌 :ok_hand: 👌🏻 :ok_hand_tone1:
//👌🏼 :ok_hand_tone2: 👌🏽 :ok_hand_tone3: 👌🏾 :ok_hand_tone4: 👌🏿 :ok_hand_tone5: 👍 :thumbsup: 👍🏻 :thumbsup_tone1: 👍🏼 :thumbsup_tone2: 👍🏽 :thumbsup_tone3: 👍🏾 :thumbsup_tone4:
//👍🏿 :thumbsup_tone5: 👎 :thumbsdown: 👎🏻 :thumbsdown_tone1: 👎🏼 :thumbsdown_tone2: 👎🏽 :thumbsdown_tone3: 👎🏾 :thumbsdown_tone4: 👎🏿 :thumbsdown_tone5: ✊ :fist: ✊🏻 :fist_tone1:
//✊🏼 :fist_tone2: ✊🏽 :fist_tone3: ✊🏾 :fist_tone4: ✊🏿 :fist_tone5: 👊 :punch: 👊🏻 :punch_tone1: 👊🏼 :punch_tone2: 👊🏽 :punch_tone3: 👊🏾 :punch_tone4:
//👊🏿 :punch_tone5: 🤛 :left_facing_fist: 🤛🏻 :left_facing_fist_tone1: 🤛🏼 :left_facing_fist_tone2: 🤛🏽 :left_facing_fist_tone3: 🤛🏾 :left_facing_fist_tone4: 🤛🏿 :left_facing_fist_tone5: 🤜 :right_facing_fist: 🤜🏻 :right_facing_fist_tone1:
//🤜🏼 :right_facing_fist_tone2: 🤜🏽 :right_facing_fist_tone3: 🤜🏾 :right_facing_fist_tone4: 🤜🏿 :right_facing_fist_tone5: 🤚 :raised_back_of_hand: 🤚🏻 :raised_back_of_hand_tone1: 🤚🏼 :raised_back_of_hand_tone2: 🤚🏽 :raised_back_of_hand_tone3: 🤚🏾 :raised_back_of_hand_tone4:
//🤚🏿 :raised_back_of_hand_tone5: 👋 :wave: 👋🏻 :wave_tone1: 👋🏼 :wave_tone2: 👋🏽 :wave_tone3: 👋🏾 :wave_tone4: 👋🏿 :wave_tone5: 🤟 :love_you_gesture: 🤟🏻 :love_you_gesture_tone1:
//🤟🏼 :love_you_gesture_tone2: 🤟🏽 :love_you_gesture_tone3: 🤟🏾 :love_you_gesture_tone4: 🤟🏿 :love_you_gesture_tone5: ✍ :writing_hand: ✍🏻 :writing_hand_tone1: ✍🏼 :writing_hand_tone2: ✍🏽 :writing_hand_tone3: ✍🏾 :writing_hand_tone4:
//✍🏿 :writing_hand_tone5: 👏 :clap: 👏🏻 :clap_tone1: 👏🏼 :clap_tone2: 👏🏽 :clap_tone3: 👏🏾 :clap_tone4: 👏🏿 :clap_tone5: 👐 :open_hands: 👐🏻 :open_hands_tone1:
//👐🏼 :open_hands_tone2: 👐🏽 :open_hands_tone3: 👐🏾 :open_hands_tone4: 👐🏿 :open_hands_tone5: 🙌 :raised_hands: 🙌🏻 :raised_hands_tone1: 🙌🏼 :raised_hands_tone2: 🙌🏽 :raised_hands_tone3: 🙌🏾 :raised_hands_tone4:
//🙌🏿 :raised_hands_tone5: 🤲 :palms_up_together: 🤲🏻 :palms_up_together_tone1: 🤲🏼 :palms_up_together_tone2: 🤲🏽 :palms_up_together_tone3: 🤲🏾 :palms_up_together_tone4: 🤲🏿 :palms_up_together_tone5: 🙏 :pray: 🙏🏻 :pray_tone1:
//🙏🏼 :pray_tone2: 🙏🏽 :pray_tone3: 🙏🏾 :pray_tone4: 🙏🏿 :pray_tone5: 🤝 :handshake: 💅 :nail_care: 💅🏻 :nail_care_tone1: 💅🏼 :nail_care_tone2: 💅🏽 :nail_care_tone3:
//💅🏾 :nail_care_tone4: 💅🏿 :nail_care_tone5: 👂 :ear: 👂🏻 :ear_tone1: 👂🏼 :ear_tone2: 👂🏽 :ear_tone3: 👂🏾 :ear_tone4: 👂🏿 :ear_tone5: 👃 :nose:
//👃🏻 :nose_tone1: 👃🏼 :nose_tone2: 👃🏽 :nose_tone3: 👃🏾 :nose_tone4: 👃🏿 :nose_tone5: 👣 :footprints: 👀 :eyes: 👁 :eye: 👁️‍🗨️ :eye_in_speech_bubble:
//🧠 :brain: 👅 :tongue: 👄 :lips: 💋 :kiss: 💘 :cupid: ❤ :heart: 💓 :heartbeat: 💔 :broken_heart: 💕 :two_hearts:
//💖 :sparkling_heart: 💗 :heartpulse: 💙 :blue_heart: 💚 :green_heart: 💛 :yellow_heart: 🧡 :orange_heart: 💜 :purple_heart: 🖤 :black_heart: 💝 :gift_heart:
//💞 :revolving_hearts: 💟 :heart_decoration: ❣ :heart_exclamation: 💌 :love_letter: 💤 :zzz: 💢 :anger: 💣 :bomb: 💥 :boom: 💦 :sweat_drops:
//💨 :dash: 💫 :dizzy: 💬 :speech_balloon: 🗨 :speech_left: 🗯 :anger_right: 💭 :thought_balloon: 🕳 :hole: 👓 :eyeglasses: 🕶 :dark_sunglasses:
//👔 :necktie: 👕 :shirt: 👖 :jeans: 🧣 :scarf: 🧤 :gloves: 🧥 :coat: 🧦 :socks: 👗 :dress: 👘 :kimono:
//👙 :bikini: 👚 :womans_clothes: 👛 :purse: 👜 :handbag: 👝 :pouch: 🛍 :shopping_bags: 🎒 :school_satchel: 👞 :mans_shoe: 👟 :athletic_shoe:
//👠 :high_heel: 👡 :sandal: 👢 :boot: 👑 :crown: 👒 :womans_hat: 🎩 :tophat: 🎓 :mortar_board: 🧢 :billed_cap: ⛑ :helmet_with_cross:
//📿 :prayer_beads: 💄 :lipstick: 💍 :ring: 💎 :gem: 🐵 :monkey_face: 🐒 :monkey: 🦍 :gorilla: 🐶 :dog: 🐕 :dog2:
//🐩 :poodle: 🐺 :wolf: 🦊 :fox: 🐱 :cat: 🐈 :cat2: 🦁 :lion_face: 🐯 :tiger: 🐅 :tiger2: 🐆 :leopard:
//🐴 :horse: 🐎 :racehorse: 🦄 :unicorn: 🦓 :zebra: 🦌 :deer: 🐮 :cow: 🐂 :ox: 🐃 :water_buffalo: 🐄 :cow2:
//🐷 :pig: 🐖 :pig2: 🐗 :boar: 🐽 :pig_nose: 🐏 :ram: 🐑 :sheep: 🐐 :goat: 🐪 :dromedary_camel: 🐫 :camel:
//🦒 :giraffe: 🐘 :elephant: 🦏 :rhino: 🐭 :mouse: 🐁 :mouse2: 🐀 :rat: 🐹 :hamster: 🐰 :rabbit: 🐇 :rabbit2:
//🐿 :chipmunk: 🦔 :hedgehog: 🦇 :bat: 🐻 :bear: 🐨 :koala: 🐼 :panda_face: 🐾 :feet: 🦃 :turkey: 🐔 :chicken:
//🐓 :rooster: 🐣 :hatching_chick: 🐤 :baby_chick: 🐥 :hatched_chick: 🐦 :bird: 🐧 :penguin: 🕊 :dove: 🦅 :eagle: 🦆 :duck:
//🦉 :owl: 🐸 :frog: 🐊 :crocodile: 🐢 :turtle: 🦎 :lizard: 🐍 :snake: 🐲 :dragon_face: 🐉 :dragon: 🦕 :sauropod:
//🦖 :t_rex: 🐳 :whale: 🐋 :whale2: 🐬 :dolphin: 🐟 :fish: 🐠 :tropical_fish: 🐡 :blowfish: 🦈 :shark: 🐙 :octopus:
//🐚 :shell: 🦀 :crab: 🦐 :shrimp: 🦑 :squid: 🐌 :snail: 🦋 :butterfly: 🐛 :bug: 🐜 :ant: 🐝 :bee:
//🐞 :beetle: 🦗 :cricket: 🕷 :spider: 🕸 :spider_web: 🦂 :scorpion: 💐 :bouquet: 🌸 :cherry_blossom: 💮 :white_flower: 🏵 :rosette:
//🌹 :rose: 🥀 :wilted_rose: 🌺 :hibiscus: 🌻 :sunflower: 🌼 :blossom: 🌷 :tulip: 🌱 :seedling: 🌲 :evergreen_tree: 🌳 :deciduous_tree:
//🌴 :palm_tree: 🌵 :cactus: 🌾 :ear_of_rice: 🌿 :herb: ☘ :shamrock: 🍀 :four_leaf_clover: 🍁 :maple_leaf: 🍂 :fallen_leaf: 🍃 :leaves:
//🍇 :grapes: 🍈 :melon: 🍉 :watermelon: 🍊 :tangerine: 🍋 :lemon: 🍌 :banana: 🍍 :pineapple: 🍎 :apple: 🍏 :green_apple:
//🍐 :pear: 🍑 :peach: 🍒 :cherries: 🍓 :strawberry: 🥝 :kiwi: 🍅 :tomato: 🥥 :coconut: 🥑 :avocado: 🍆 :eggplant:
//🥔 :potato: 🥕 :carrot: 🌽 :corn: 🌶 :hot_pepper: 🥒 :cucumber: 🥦 :broccoli: 🍄 :mushroom: 🥜 :peanuts: 🌰 :chestnut:
//🍞 :bread: 🥐 :croissant: 🥖 :french_bread: 🥨 :pretzel: 🥞 :pancakes: 🧀 :cheese: 🍖 :meat_on_bone: 🍗 :poultry_leg: 🥩 :cut_of_meat:
//🥓 :bacon: 🍔 :hamburger: 🍟 :fries: 🍕 :pizza: 🌭 :hotdog: 🥪 :sandwich: 🌮 :taco: 🌯 :burrito: 🥙 :stuffed_flatbread:
//🥚 :egg: 🍳 :cooking: 🥘 :shallow_pan_of_food: 🍲 :stew: 🥣 :bowl_with_spoon: 🥗 :salad: 🍿 :popcorn: 🥫 :canned_food: 🍱 :bento:
//🍘 :rice_cracker: 🍙 :rice_ball: 🍚 :rice: 🍛 :curry: 🍜 :ramen: 🍝 :spaghetti: 🍠 :sweet_potato: 🍢 :oden: 🍣 :sushi:
//🍤 :fried_shrimp: 🍥 :fish_cake: 🍡 :dango: 🥟 :dumpling: 🥠 :fortune_cookie: 🥡 :takeout_box: 🍦 :icecream: 🍧 :shaved_ice: 🍨 :ice_cream:
//🍩 :doughnut: 🍪 :cookie: 🎂 :birthday: 🍰 :cake: 🥧 :pie: 🍫 :chocolate_bar: 🍬 :candy: 🍭 :lollipop: 🍮 :custard:
//🍯 :honey_pot: 🍼 :baby_bottle: 🥛 :milk: ☕ :coffee: 🍵 :tea: 🍶 :sake: 🍾 :champagne: 🍷 :wine_glass: 🍸 :cocktail:
//🍹 :tropical_drink: 🍺 :beer: 🍻 :beers: 🥂 :champagne_glass: 🥃 :tumbler_glass: 🥤 :cup_with_straw: 🥢 :chopsticks: 🍽 :fork_knife_plate: 🍴 :fork_and_knife:
//🥄 :spoon: 🔪 :knife: 🏺 :amphora: 🌍 :earth_africa: 🌎 :earth_americas: 🌏 :earth_asia: 🌐 :globe_with_meridians: 🗺 :map: 🗾 :japan:
//🏔 :mountain_snow: ⛰ :mountain: 🌋 :volcano: 🗻 :mount_fuji: 🏕 :camping: 🏖 :beach: 🏜 :desert: 🏝 :island: 🏞 :park:
//🏟 :stadium: 🏛 :classical_building: 🏗 :construction_site: 🏘 :homes: 🏙 :cityscape: 🏚 :house_abandoned: 🏠 :house: 🏡 :house_with_garden: 🏢 :office:
//🏣 :post_office: 🏤 :european_post_office: 🏥 :hospital: 🏦 :bank: 🏨 :hotel: 🏩 :love_hotel: 🏪 :convenience_store: 🏫 :school: 🏬 :department_store:
//🏭 :factory: 🏯 :japanese_castle: 🏰 :european_castle: 💒 :wedding: 🗼 :tokyo_tower: 🗽 :statue_of_liberty: ⛪ :church: 🕌 :mosque: 🕍 :synagogue:
//⛩ :shinto_shrine: 🕋 :kaaba: ⛲ :fountain: ⛺ :tent: 🌁 :foggy: 🌃 :night_with_stars: 🌄 :sunrise_over_mountains: 🌅 :sunrise: 🌆 :city_dusk:
//🌇 :city_sunset: 🌉 :bridge_at_night: ♨ :hotsprings: 🌌 :milky_way: 🎠 :carousel_horse: 🎡 :ferris_wheel: 🎢 :roller_coaster: 💈 :barber: 🎪 :circus_tent:
//🎭 :performing_arts: 🖼 :frame_photo: 🎨 :art: 🎰 :slot_machine: 🚂 :steam_locomotive: 🚃 :railway_car: 🚄 :bullettrain_side: 🚅 :bullettrain_front: 🚆 :train2:
//🚇 :metro: 🚈 :light_rail: 🚉 :station: 🚊 :tram: 🚝 :monorail: 🚞 :mountain_railway: 🚋 :train: 🚌 :bus: 🚍 :oncoming_bus:
//🚎 :trolleybus: 🚐 :minibus: 🚑 :ambulance: 🚒 :fire_engine: 🚓 :police_car: 🚔 :oncoming_police_car: 🚕 :taxi: 🚖 :oncoming_taxi: 🚗 :red_car:
//🚘 :oncoming_automobile: 🚙 :blue_car: 🚚 :truck: 🚛 :articulated_lorry: 🚜 :tractor: 🚲 :bike: 🛴 :scooter: 🛵 :motor_scooter: 🚏 :busstop:
//🛣 :motorway: 🛤 :railway_track: ⛽ :fuelpump: 🚨 :rotating_light: 🚥 :traffic_light: 🚦 :vertical_traffic_light: 🚧 :construction: 🛑 :octagonal_sign: ⚓ :anchor:
//⛵ :sailboat: 🛶 :canoe: 🚤 :speedboat: 🛳 :cruise_ship: ⛴ :ferry: 🛥 :motorboat: 🚢 :ship: ✈ :airplane: 🛩 :airplane_small:
//🛫 :airplane_departure: 🛬 :airplane_arriving: 💺 :seat: 🚁 :helicopter: 🚟 :suspension_railway: 🚠 :mountain_cableway: 🚡 :aerial_tramway: 🛰 :satellite_orbital: 🚀 :rocket:
//🛸 :flying_saucer: 🛎 :bellhop: 🚪 :door: 🛏 :bed: 🛋 :couch: 🚽 :toilet: 🚿 :shower: 🛁 :bathtub: ⌛ :hourglass:
//⏳ :hourglass_flowing_sand: ⌚ :watch: ⏰ :alarm_clock: ⏱ :stopwatch: ⏲ :timer: 🕰 :clock: 🕛 :clock12: 🕧 :clock1230: 🕐 :clock1:
//🕜 :clock130: 🕑 :clock2: 🕝 :clock230: 🕒 :clock3: 🕞 :clock330: 🕓 :clock4: 🕟 :clock430: 🕔 :clock5: 🕠 :clock530:
//🕕 :clock6: 🕡 :clock630: 🕖 :clock7: 🕢 :clock730: 🕗 :clock8: 🕣 :clock830: 🕘 :clock9: 🕤 :clock930: 🕙 :clock10:
//🕥 :clock1030: 🕚 :clock11: 🕦 :clock1130: 🌑 :new_moon: 🌒 :waxing_crescent_moon: 🌓 :first_quarter_moon: 🌔 :waxing_gibbous_moon: 🌕 :full_moon: 🌖 :waning_gibbous_moon:
//🌗 :last_quarter_moon: 🌘 :waning_crescent_moon: 🌙 :crescent_moon: 🌚 :new_moon_with_face: 🌛 :first_quarter_moon_with_face: 🌜 :last_quarter_moon_with_face: 🌡 :thermometer: ☀ :sunny: 🌝 :full_moon_with_face:
//🌞 :sun_with_face: ⭐ :star: 🌟 :star2: 🌠 :stars: ☁ :cloud: ⛅ :partly_sunny: ⛈ :thunder_cloud_rain: 🌤 :white_sun_small_cloud: 🌥 :white_sun_cloud:
//🌦 :white_sun_rain_cloud: 🌧 :cloud_rain: 🌨 :cloud_snow: 🌩 :cloud_lightning: 🌪 :cloud_tornado: 🌫 :fog: 🌬 :wind_blowing_face: 🌀 :cyclone: 🌈 :rainbow:
//🌂 :closed_umbrella: ☂ :umbrella2: ☔ :umbrella: ⛱ :beach_umbrella: ⚡ :zap: ❄ :snowflake: ☃ :snowman2: ⛄ :snowman: ☄ :comet:
//🔥 :fire: 💧 :droplet: 🌊 :ocean: 🎃 :jack_o_lantern: 🎄 :christmas_tree: 🎆 :fireworks: 🎇 :sparkler: ✨ :sparkles: 🎈 :balloon:
//🎉 :tada: 🎊 :confetti_ball: 🎋 :tanabata_tree: 🎍 :bamboo: 🎎 :dolls: 🎏 :flags: 🎐 :wind_chime: 🎑 :rice_scene: 🎀 :ribbon:
//🎁 :gift: 🎗 :reminder_ribbon: 🎟 :tickets: 🎫 :ticket: 🎖 :military_medal: 🏆 :trophy: 🏅 :medal: 🥇 :first_place: 🥈 :second_place:
//🥉 :third_place: ⚽ :soccer: ⚾ :baseball: 🏀 :basketball: 🏐 :volleyball: 🏈 :football: 🏉 :rugby_football: 🎾 :tennis: 🎱 :8ball:
//🎳 :bowling: 🏏 :cricket_game: 🏑 :field_hockey: 🏒 :hockey: 🏓 :ping_pong: 🏸 :badminton: 🥊 :boxing_glove: 🥋 :martial_arts_uniform: 🥅 :goal:
//🎯 :dart: ⛳ :golf: ⛸ :ice_skate: 🎣 :fishing_pole_and_fish: 🎽 :running_shirt_with_sash: 🎿 :ski: 🛷 :sled: 🥌 :curling_stone: 🎮 :video_game:
//🕹 :joystick: 🎲 :game_die: ♠ :spades: ♥ :hearts: ♦ :diamonds: ♣ :clubs: 🃏 :black_joker: 🀄 :mahjong: 🎴 :flower_playing_cards:
//🔇 :mute: 🔈 :speaker: 🔉 :sound: 🔊 :loud_sound: 📢 :loudspeaker: 📣 :mega: 📯 :postal_horn: 🔔 :bell: 🔕 :no_bell:
//🎼 :musical_score: 🎵 :musical_note: 🎶 :notes: 🎙 :microphone2: 🎚 :level_slider: 🎛 :control_knobs: 🎤 :microphone: 🎧 :headphones: 📻 :radio:
//🎷 :saxophone: 🎸 :guitar: 🎹 :musical_keyboard: 🎺 :trumpet: 🎻 :violin: 🥁 :drum: 📱 :iphone: 📲 :calling: ☎ :telephone:
//📞 :telephone_receiver: 📟 :pager: 📠 :fax: 🔋 :battery: 🔌 :electric_plug: 💻 :computer: 🖥 :desktop: 🖨 :printer: ⌨ :keyboard:
//🖱 :mouse_three_button: 🖲 :trackball: 💽 :minidisc: 💾 :floppy_disk: 💿 :cd: 📀 :dvd: 🎥 :movie_camera: 🎞 :film_frames: 📽 :projector:
//🎬 :clapper: 📺 :tv: 📷 :camera: 📸 :camera_with_flash: 📹 :video_camera: 📼 :vhs: 🔍 :mag: 🔎 :mag_right: 🔬 :microscope:
//🔭 :telescope: 📡 :satellite: 🕯 :candle: 💡 :bulb: 🔦 :flashlight: 🏮 :izakaya_lantern: 📔 :notebook_with_decorative_cover: 📕 :closed_book: 📖 :book:
//📗 :green_book: 📘 :blue_book: 📙 :orange_book: 📚 :books: 📓 :notebook: 📒 :ledger: 📃 :page_with_curl: 📜 :scroll: 📄 :page_facing_up:
//📰 :newspaper: 🗞 :newspaper2: 📑 :bookmark_tabs: 🔖 :bookmark: 🏷 :label: 💰 :moneybag: 💴 :yen: 💵 :dollar: 💶 :euro:
//💷 :pound: 💸 :money_with_wings: 💳 :credit_card: 💹 :chart: 💱 :currency_exchange: 💲 :heavy_dollar_sign: ✉ :envelope: 📧 :e-mail: 📨 :incoming_envelope:
//📩 :envelope_with_arrow: 📤 :outbox_tray: 📥 :inbox_tray: 📦 :package: 📫 :mailbox: 📪 :mailbox_closed: 📬 :mailbox_with_mail: 📭 :mailbox_with_no_mail: 📮 :postbox:
//🗳 :ballot_box: ✏ :pencil2: ✒ :black_nib: 🖋 :pen_fountain: 🖊 :pen_ballpoint: 🖌 :paintbrush: 🖍 :crayon: 📝 :pencil: 💼 :briefcase:
//📁 :file_folder: 📂 :open_file_folder: 🗂 :dividers: 📅 :date: 📆 :calendar: 🗒 :notepad_spiral: 🗓 :calendar_spiral: 📇 :card_index: 📈 :chart_with_upwards_trend:
//📉 :chart_with_downwards_trend: 📊 :bar_chart: 📋 :clipboard: 📌 :pushpin: 📍 :round_pushpin: 📎 :paperclip: 🖇 :paperclips: 📏 :straight_ruler: 📐 :triangular_ruler:
//✂ :scissors: 🗃 :card_box: 🗄 :file_cabinet: 🗑 :wastebasket: 🔒 :lock: 🔓 :unlock: 🔏 :lock_with_ink_pen: 🔐 :closed_lock_with_key: 🔑 :key:
//🗝 :key2: 🔨 :hammer: ⛏ :pick: ⚒ :hammer_pick: 🛠 :tools: 🗡 :dagger: ⚔ :crossed_swords: 🔫 :gun: 🏹 :bow_and_arrow:
//🛡 :shield: 🔧 :wrench: 🔩 :nut_and_bolt: ⚙ :gear: 🗜 :compression: ⚗ :alembic: ⚖ :scales: 🔗 :link: ⛓ :chains:
//💉 :syringe: 💊 :pill: 🚬 :smoking: ⚰ :coffin: ⚱ :urn: 🗿 :moyai: 🛢 :oil: 🔮 :crystal_ball: 🛒 :shopping_cart:
//🏧 :atm: 🚮 :put_litter_in_its_place: 🚰 :potable_water: ♿ :wheelchair: 🚹 :mens: 🚺 :womens: 🚻 :restroom: 🚼 :baby_symbol: 🚾 :wc:
//🛂 :passport_control: 🛃 :customs: 🛄 :baggage_claim: 🛅 :left_luggage: ⚠ :warning: 🚸 :children_crossing: ⛔ :no_entry: 🚫 :no_entry_sign: 🚳 :no_bicycles:
//🚭 :no_smoking: 🚯 :do_not_litter: 🚱 :non-potable_water: 🚷 :no_pedestrians: 📵 :no_mobile_phones: 🔞 :underage: ☢ :radioactive: ☣ :biohazard: ⬆ :arrow_up:
//↗ :arrow_upper_right: ➡ :arrow_right: ↘ :arrow_lower_right: ⬇ :arrow_down: ↙ :arrow_lower_left: ⬅ :arrow_left: ↖ :arrow_upper_left: ↕ :arrow_up_down: ↔ :left_right_arrow:
//↩ :leftwards_arrow_with_hook: ↪ :arrow_right_hook: ⤴ :arrow_heading_up: ⤵ :arrow_heading_down: 🔃 :arrows_clockwise: 🔄 :arrows_counterclockwise: 🔙 :back: 🔚 :end: 🔛 :on:
//🔜 :soon: 🔝 :top: 🛐 :place_of_worship: ⚛ :atom: 🕉 :om_symbol: ✡ :star_of_david: ☸ :wheel_of_dharma: ☯ :yin_yang: ✝ :cross:
//☦ :orthodox_cross: ☪ :star_and_crescent: ☮ :peace: 🕎 :menorah: 🔯 :six_pointed_star: ♈ :aries: ♉ :taurus: ♊ :gemini: ♋ :cancer:
//♌ :leo: ♍ :virgo: ♎ :libra: ♏ :scorpius: ♐ :sagittarius: ♑ :capricorn: ♒ :aquarius: ♓ :pisces: ⛎ :ophiuchus:
//🔀 :twisted_rightwards_arrows: 🔁 :repeat: 🔂 :repeat_one: ▶ :arrow_forward: ⏩ :fast_forward: ⏭ :track_next: ⏯ :play_pause: ◀ :arrow_backward: ⏪ :rewind:
//⏮ :track_previous: 🔼 :arrow_up_small: ⏫ :arrow_double_up: 🔽 :arrow_down_small: ⏬ :arrow_double_down: ⏸ :pause_button: ⏹ :stop_button: ⏺ :record_button: ⏏ :eject:
//🎦 :cinema: 🔅 :low_brightness: 🔆 :high_brightness: 📶 :signal_strength: 📳 :vibration_mode: 📴 :mobile_phone_off: ♀ :female_sign: ♂ :male_sign: ⚕ :medical_symbol:
//♻ :recycle: ⚜ :fleur-de-lis: 🔱 :trident: 📛 :name_badge: 🔰 :beginner: ⭕ :o: ✅ :white_check_mark: ☑ :ballot_box_with_check: ✔ :heavy_check_mark:
//✖ :heavy_multiplication_x: ❌ :x: ❎ :negative_squared_cross_mark: ➕ :heavy_plus_sign: ➖ :heavy_minus_sign: ➗ :heavy_division_sign: ➰ :curly_loop: ➿ :loop: 〽 :part_alternation_mark:
//✳ :eight_spoked_asterisk: ✴ :eight_pointed_black_star: ❇ :sparkle: ‼ :bangbang: ⁉ :interrobang: ❓ :question: ❔ :grey_question: ❕ :grey_exclamation: ❗ :exclamation:
//〰 :wavy_dash: © :copyright: ® :registered: ™ :tm: #️⃣ :hash: *️⃣ :asterisk: 0️⃣ :zero: 1️⃣ :one: 2️⃣ :two:
//3️⃣ :three: 4️⃣ :four: 5️⃣ :five: 6️⃣ :six: 7️⃣ :seven: 8️⃣ :eight: 9️⃣ :nine: 🔟 :keycap_ten: 💯 :100:
//🔠 :capital_abcd: 🔡 :abcd: 🔢 :1234: 🔣 :symbols: 🔤 :abc: 🅰 :a: 🆎 :ab: 🅱 :b: 🆑 :cl:
//🆒 :cool: 🆓 :free: ℹ :information_source: 🆔 :id: Ⓜ :m: 🆕 :new: 🆖 :ng: 🅾 :o2: 🆗 :ok:
//🅿 :parking: 🆘 :sos: 🆙 :up: 🆚 :vs: 🈁 :koko: 🈂 :sa: 🈷 :u6708: 🈶 :u6709: 🈯 :u6307:
//🉐 :ideograph_advantage: 🈹 :u5272: 🈚 :u7121: 🈲 :u7981: 🉑 :accept: 🈸 :u7533: 🈴 :u5408: 🈳 :u7a7a: ㊗ :congratulations:
//㊙ :secret: 🈺 :u55b6: 🈵 :u6e80: ▪ :black_small_square: ▫ :white_small_square: ◻ :white_medium_square: ◼ :black_medium_square: ◽ :white_medium_small_square: ◾ :black_medium_small_square:
//⬛ :black_large_square: ⬜ :white_large_square: 🔶 :large_orange_diamond: 🔷 :large_blue_diamond: 🔸 :small_orange_diamond: 🔹 :small_blue_diamond: 🔺 :small_red_triangle: 🔻 :small_red_triangle_down: 💠 :diamond_shape_with_a_dot_inside:
//🔘 :radio_button: 🔲 :black_square_button: 🔳 :white_square_button: ⚪ :white_circle: ⚫ :black_circle: 🔴 :red_circle: 🔵 :blue_circle: 🏁 :checkered_flag: 🚩 :triangular_flag_on_post:
//🎌 :crossed_flags: 🏴 :flag_black: 🏳 :flag_white: 🏳️‍🌈 :rainbow_flag: 🇦🇨 :flag_ac: 🇦🇩 :flag_ad: 🇦🇪 :flag_ae: 🇦🇫 :flag_af: 🇦🇬 :flag_ag:
//🇦🇮 :flag_ai: 🇦🇱 :flag_al: 🇦🇲 :flag_am: 🇦🇴 :flag_ao: 🇦🇶 :flag_aq: 🇦🇷 :flag_ar: 🇦🇸 :flag_as: 🇦🇹 :flag_at: 🇦🇺 :flag_au:
//🇦🇼 :flag_aw: 🇦🇽 :flag_ax: 🇦🇿 :flag_az: 🇧🇦 :flag_ba: 🇧🇧 :flag_bb: 🇧🇩 :flag_bd: 🇧🇪 :flag_be: 🇧🇫 :flag_bf: 🇧🇬 :flag_bg:
//🇧🇭 :flag_bh: 🇧🇮 :flag_bi: 🇧🇯 :flag_bj: 🇧🇱 :flag_bl: 🇧🇲 :flag_bm: 🇧🇳 :flag_bn: 🇧🇴 :flag_bo: 🇧🇶 :flag_bq: 🇧🇷 :flag_br:
//🇧🇸 :flag_bs: 🇧🇹 :flag_bt: 🇧🇻 :flag_bv: 🇧🇼 :flag_bw: 🇧🇾 :flag_by: 🇧🇿 :flag_bz: 🇨🇦 :flag_ca: 🇨🇨 :flag_cc: 🇨🇩 :flag_cd:
//🇨🇫 :flag_cf: 🇨🇬 :flag_cg: 🇨🇭 :flag_ch: 🇨🇮 :flag_ci: 🇨🇰 :flag_ck: 🇨🇱 :flag_cl: 🇨🇲 :flag_cm: 🇨🇳 :flag_cn: 🇨🇴 :flag_co:
//🇨🇵 :flag_cp: 🇨🇷 :flag_cr: 🇨🇺 :flag_cu: 🇨🇻 :flag_cv: 🇨🇼 :flag_cw: 🇨🇽 :flag_cx: 🇨🇾 :flag_cy: 🇨🇿 :flag_cz: 🇩🇪 :flag_de:
//🇩🇬 :flag_dg: 🇩🇯 :flag_dj: 🇩🇰 :flag_dk: 🇩🇲 :flag_dm: 🇩🇴 :flag_do: 🇩🇿 :flag_dz: 🇪🇦 :flag_ea: 🇪🇨 :flag_ec: 🇪🇪 :flag_ee:
//🇪🇬 :flag_eg: 🇪🇭 :flag_eh: 🇪🇷 :flag_er: 🇪🇸 :flag_es: 🇪🇹 :flag_et: 🇪🇺 :flag_eu: 🇫🇮 :flag_fi: 🇫🇯 :flag_fj: 🇫🇰 :flag_fk:
//🇫🇲 :flag_fm: 🇫🇴 :flag_fo: 🇫🇷 :flag_fr: 🇬🇦 :flag_ga: 🇬🇧 :flag_gb: 🇬🇩 :flag_gd: 🇬🇪 :flag_ge: 🇬🇫 :flag_gf: 🇬🇬 :flag_gg:
//🇬🇭 :flag_gh: 🇬🇮 :flag_gi: 🇬🇱 :flag_gl: 🇬🇲 :flag_gm: 🇬🇳 :flag_gn: 🇬🇵 :flag_gp: 🇬🇶 :flag_gq: 🇬🇷 :flag_gr: 🇬🇸 :flag_gs:
//🇬🇹 :flag_gt: 🇬🇺 :flag_gu: 🇬🇼 :flag_gw: 🇬🇾 :flag_gy: 🇭🇰 :flag_hk: 🇭🇲 :flag_hm: 🇭🇳 :flag_hn: 🇭🇷 :flag_hr: 🇭🇹 :flag_ht:
//🇭🇺 :flag_hu: 🇮🇨 :flag_ic: 🇮🇩 :flag_id: 🇮🇪 :flag_ie: 🇮🇱 :flag_il: 🇮🇲 :flag_im: 🇮🇳 :flag_in: 🇮🇴 :flag_io: 🇮🇶 :flag_iq:
//🇮🇷 :flag_ir: 🇮🇸 :flag_is: 🇮🇹 :flag_it: 🇯🇪 :flag_je: 🇯🇲 :flag_jm: 🇯🇴 :flag_jo: 🇯🇵 :flag_jp: 🇰🇪 :flag_ke: 🇰🇬 :flag_kg:
//🇰🇭 :flag_kh: 🇰🇮 :flag_ki: 🇰🇲 :flag_km: 🇰🇳 :flag_kn: 🇰🇵 :flag_kp: 🇰🇷 :flag_kr: 🇰🇼 :flag_kw: 🇰🇾 :flag_ky: 🇰🇿 :flag_kz:
//🇱🇦 :flag_la: 🇱🇧 :flag_lb: 🇱🇨 :flag_lc: 🇱🇮 :flag_li: 🇱🇰 :flag_lk: 🇱🇷 :flag_lr: 🇱🇸 :flag_ls: 🇱🇹 :flag_lt: 🇱🇺 :flag_lu:
//🇱🇻 :flag_lv: 🇱🇾 :flag_ly: 🇲🇦 :flag_ma: 🇲🇨 :flag_mc: 🇲🇩 :flag_md: 🇲🇪 :flag_me: 🇲🇫 :flag_mf: 🇲🇬 :flag_mg: 🇲🇭 :flag_mh:
//🇲🇰 :flag_mk: 🇲🇱 :flag_ml: 🇲🇲 :flag_mm: 🇲🇳 :flag_mn: 🇲🇴 :flag_mo: 🇲🇵 :flag_mp: 🇲🇶 :flag_mq: 🇲🇷 :flag_mr: 🇲🇸 :flag_ms:
//🇲🇹 :flag_mt: 🇲🇺 :flag_mu: 🇲🇻 :flag_mv: 🇲🇼 :flag_mw: 🇲🇽 :flag_mx: 🇲🇾 :flag_my: 🇲🇿 :flag_mz: 🇳🇦 :flag_na: 🇳🇨 :flag_nc:
//🇳🇪 :flag_ne: 🇳🇫 :flag_nf: 🇳🇬 :flag_ng: 🇳🇮 :flag_ni: 🇳🇱 :flag_nl: 🇳🇴 :flag_no: 🇳🇵 :flag_np: 🇳🇷 :flag_nr: 🇳🇺 :flag_nu:
//🇳🇿 :flag_nz: 🇴🇲 :flag_om: 🇵🇦 :flag_pa: 🇵🇪 :flag_pe: 🇵🇫 :flag_pf: 🇵🇬 :flag_pg: 🇵🇭 :flag_ph: 🇵🇰 :flag_pk: 🇵🇱 :flag_pl:
//🇵🇲 :flag_pm: 🇵🇳 :flag_pn: 🇵🇷 :flag_pr: 🇵🇸 :flag_ps: 🇵🇹 :flag_pt: 🇵🇼 :flag_pw: 🇵🇾 :flag_py: 🇶🇦 :flag_qa: 🇷🇪 :flag_re:
//🇷🇴 :flag_ro: 🇷🇸 :flag_rs: 🇷🇺 :flag_ru: 🇷🇼 :flag_rw: 🇸🇦 :flag_sa: 🇸🇧 :flag_sb: 🇸🇨 :flag_sc: 🇸🇩 :flag_sd: 🇸🇪 :flag_se:
//🇸🇬 :flag_sg: 🇸🇭 :flag_sh: 🇸🇮 :flag_si: 🇸🇯 :flag_sj: 🇸🇰 :flag_sk: 🇸🇱 :flag_sl: 🇸🇲 :flag_sm: 🇸🇳 :flag_sn: 🇸🇴 :flag_so:
//🇸🇷 :flag_sr: 🇸🇸 :flag_ss: 🇸🇹 :flag_st: 🇸🇻 :flag_sv: 🇸🇽 :flag_sx: 🇸🇾 :flag_sy: 🇸🇿 :flag_sz: 🇹🇦 :flag_ta: 🇹🇨 :flag_tc:
//🇹🇩 :flag_td: 🇹🇫 :flag_tf: 🇹🇬 :flag_tg: 🇹🇭 :flag_th: 🇹🇯 :flag_tj: 🇹🇰 :flag_tk: 🇹🇱 :flag_tl: 🇹🇲 :flag_tm: 🇹🇳 :flag_tn:
//🇹🇴 :flag_to: 🇹🇷 :flag_tr: 🇹🇹 :flag_tt: 🇹🇻 :flag_tv: 🇹🇼 :flag_tw: 🇹🇿 :flag_tz: 🇺🇦 :flag_ua: 🇺🇬 :flag_ug: 🇺🇲 :flag_um:
//🇺🇳 :united_nations: 🇺🇸 :flag_us: 🇺🇾 :flag_uy: 🇺🇿 :flag_uz: 🇻🇦 :flag_va: 🇻🇨 :flag_vc: 🇻🇪 :flag_ve: 🇻🇬 :flag_vg: 🇻🇮 :flag_vi:
//🇻🇳 :flag_vn: 🇻🇺 :flag_vu: 🇼🇫 :flag_wf: 🇼🇸 :flag_ws: 🇽🇰 :flag_xk: 🇾🇪 :flag_ye: 🇾🇹 :flag_yt: 🇿🇦 :flag_za: 🇿🇲 :flag_zm:
//🇿🇼 :flag_zw: 🏴󠁧󠁢󠁥󠁮󠁧󠁿 :england: 🏴󠁧󠁢󠁳󠁣󠁴󠁿 :scotland: 🏴󠁧󠁢󠁷󠁬󠁳󠁿 :wales:

func TestEmoji(t *testing.T) {
	emoji.Println(":white_check_mark:")
	emoji.Println(":negative_squared_cross_mark:")
}

func TestEmojiJson(t *testing.T) {
	//url := "https://raw.githubusercontent.com/CodeFreezr/emojo/master/db/v5/emoji-v5.json"
	open, err := os.Open("./emoji.txt")
	if err != nil {
		return
	}
	all, err := io.ReadAll(open)
	if err != nil {
		return
	}
	type T struct {
		No          int    `json:"No"`
		Emoji       string `json:"Emoji"`
		Category    string `json:"Category"`
		SubCategory string `json:"SubCategory"`
		Unicode     string `json:"Unicode"`
		Name        string `json:"Name"`
		Tags        string `json:"Tags"`
		Shortcode   string `json:"Shortcode"`
	}
	var ts []T
	err = json.Unmarshal(all, &ts)
	if err != nil {
		return
	}
	k := 0
	for _, t2 := range ts {
		k++
		fmt.Printf("%s %s ", t2.Emoji, t2.Shortcode)
		if k > 8 {
			fmt.Println()
			k = 0
		}
	}
}
