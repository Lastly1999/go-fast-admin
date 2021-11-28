import { ref } from "vue"


const useCommonDialog = () => {
  const title = ref<string>('')
  const width = ref<string>('40%')
  const loading = ref<boolean>(false)
  const visible = ref<boolean>(false)
  return {
    title,
    width,
    loading,
    visible
  }
}

export default useCommonDialog