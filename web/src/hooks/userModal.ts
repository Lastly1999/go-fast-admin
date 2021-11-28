import { ref } from "vue"


export default function (title: string, show: boolean) {
  const modalTitle = ref(title)
  const visible = ref(show)

  const showModal = (): void => {
    visible.value = true
  }

  const closeModal = (): void => {
    visible.value = false
  }

  return {
    modalTitle,
    visible,
    showModal,
    closeModal
  }
} 